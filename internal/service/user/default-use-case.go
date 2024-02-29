package user

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bowoBp/myDate/internal/adapter/repository"
	"github.com/bowoBp/myDate/internal/constant"
	"github.com/bowoBp/myDate/internal/domains"
	global_usage "github.com/bowoBp/myDate/internal/service/global-usage"
	"github.com/bowoBp/myDate/pkg/environment"
	"github.com/bowoBp/myDate/pkg/mailing"
	"github.com/bowoBp/myDate/pkg/maker"
	"github.com/bowoBp/myDate/pkg/middleware"
	time2 "github.com/bowoBp/myDate/pkg/time"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	UseCase struct {
		userRepo    repository.UserRepoInterface
		maker       maker.Generator
		env         environment.Environment
		clock       time2.Clock
		globalUsage global_usage.UsecaseInterface
		smtp        mailing.SendInBlueInterface
		auth        middleware.Auth
	}

	UseCaseInterface interface {
		AddUser(
			ctx context.Context,
			payload RegisterPayload,
		) (result RegisterResult, err error)
		VerifyOtp(ctx context.Context, payload VerifyOtpPayload) (bool, error)
		Login(
			ctx context.Context,
			payload LoginPayload,
		) (domains.User, string, error)
	}
)

func (uc UseCase) AddUser(
	ctx context.Context,
	payload RegisterPayload,
) (result RegisterResult, err error) {
	res, err := uc.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil || res.Email == payload.Email {
		log.Println(err)
		return RegisterResult{}, err
	}

	user, err := uc.registerUser(ctx, payload)

	result.User = RegisterResponseData{
		ID:       user.UserID,
		Name:     user.Username,
		Email:    user.Email,
		IsActive: user.IsActive,
	}

	otpResult, err := uc.GenerateAndSendOTP(
		ctx,
		SendOtpPayload{
			Email:  payload.Email,
			Name:   payload.Name,
			UserID: uint64(user.UserID),
		},
		false,
	)
	if err != nil {
		log.Println(err)
		return RegisterResult{}, err
	}

	result.Email = otpResult.EmailResult

	user.Otp = otpResult.Otp
	_, err = uc.userRepo.UpdateSelectedField(ctx, user, "otp")
	if err != nil {
		log.Println(err)
		return RegisterResult{}, fmt.Errorf("uc.userRepo.UpdateSelecteField: %w", err)
	}
	return result, nil
}

func (uc UseCase) VerifyOtp(ctx context.Context, payload VerifyOtpPayload) (bool, error) {
	user, err := uc.userRepo.GetUserByID(ctx, uint(payload.UserID))
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, constant.ErrUserNameNotFound
		}
		return false, fmt.Errorf("uc.userRepo.GetUserByID: %w", err)
	}

	otp, err := strconv.ParseInt(payload.Otp, 10, 32)
	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("strconv.ParseInt: %w", err)
	}

	if user.IsActive {
		return true, nil
	}

	switch user.Otp {
	case 0:
		return false, constant.ErrNewOtpRequired
	case int(otp):
		user.IsActive = true
		_, err = uc.userRepo.UpdateSelectedField(ctx, user, "is_active")
		if err != nil {
			log.Println(err)
			return false, fmt.Errorf("userRepo.UpdateSelecteField: %w", err)
		}
		return true, nil
	default:
		user.Otp = 0
		_, err := uc.userRepo.UpdateSelectedField(ctx, user, "OTPCode")
		if err != nil {
			log.Println(err)
			return false, fmt.Errorf("userRepo.UpdateSelecteField: %w", err)
		}
		return false, constant.ErrOtpInvalid
	}
}

func (uc UseCase) registerUser(
	ctx context.Context,
	payload RegisterPayload,
) (result *domains.User, err error) {
	encryptPwd, err := uc.maker.EncryptMessage(
		[]byte(uc.env.Get("SECRET_ENCRYPTION_PASS")),
		[]byte(payload.Password),
	)
	if err != nil {
		log.Println(err)
		return &domains.User{}, fmt.Errorf("uc.maker.HashAndSalt: %w", err)
	}
	user, errAdd := uc.userRepo.AddUser(
		ctx,
		&domains.User{

			Username:         payload.Name,
			Email:            payload.Email,
			Password:         hex.EncodeToString(encryptPwd),
			RegistrationDate: time.Now(),
			PremiumStatus:    false,
			Otp:              0,
			IsActive:         false,
		},
	)
	return user, errAdd
}

func (uc UseCase) GenerateAndSendOTP(
	ctx context.Context,
	emailPayload SendOtpPayload,
	regenerate bool,
) (
	result SendOtpResult,
	err error,
) {
	var user *domains.User
	if regenerate {
		user, err = uc.userRepo.GetUserByID(ctx, uint(emailPayload.UserID))
		if err != nil {
			log.Println(err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return SendOtpResult{}, constant.ErrUserNameIsUsed
			}
			return SendOtpResult{}, fmt.Errorf("uc.userRepo.GetUserByID: %w", err)
		}
		emailPayload.Email = user.Email
		emailPayload.Name = user.Username

		if user.IsActive {
			return SendOtpResult{}, constant.ErrResendOTPNotAllowed
		}
	}

	movingFactor := uint64(uc.clock.NowUnix() / 30)
	secret := uc.env.Get("HOTP_SECRET")
	otp, err := uc.maker.GenerateOTPCode(secret, movingFactor)
	if err != nil {
		log.Println(err)
		return SendOtpResult{}, fmt.Errorf("uc.davinci.GenerateOTPCode: %w", err)
	}

	otpStr := strconv.Itoa(otp)

	var tmplData = global_usage.EmailBodyVerifyOTPPayload{
		Name:       emailPayload.Name,
		OTPs:       strings.Split(otpStr, ""),
		VerifyPage: os.Getenv("FRONT_END_HOST") + "/register/verifikasi/" + strconv.Itoa(int(emailPayload.UserID)),
	}
	emailPayload.Content, err = uc.globalUsage.GenerateEmailBodyVerifyOTP(tmplData)
	emailPayload.Subject = "Konfirmasi Email Owner"
	err = uc.sendEmail(ctx, emailPayload)
	if err != nil {
		log.Println(err)
		return SendOtpResult{}, err
	}

	if regenerate {
		user.Otp = otp
		user.IsActive = false
		_, err := uc.userRepo.UpdateSelectedField(ctx, user, "otp", "is_active")
		if err != nil {
			log.Println(err)
			return SendOtpResult{}, fmt.Errorf("uc.userRepo.UpdateSelecteField: %w", err)
		}
	}

	return SendOtpResult{
		Otp: otp,
	}, nil
}

func (uc UseCase) sendEmail(ctx context.Context, emailPayload SendOtpPayload) error {
	err := uc.smtp.NativeSendEmail(ctx, mailing.NativeSendEmailPayload{
		Host:     os.Getenv("SMPT_SERVER_HOST"),
		Port:     os.Getenv("SMPT_SERVER_PORT"),
		Subject:  emailPayload.Subject,
		Username: os.Getenv("SUPPORT_EMAIL"),
		Password: os.Getenv("SUPPORT_EMAIL_PASS"),
		SendTo:   emailPayload.Email,
		HtmlBody: emailPayload.Content,
	})
	if err != nil {
		log.Println(err)
		return fmt.Errorf("uc.smtp.SendEmail: %w", err)
	}
	return nil
}

func (uc UseCase) Login(
	ctx context.Context,
	payload LoginPayload,
) (domains.User, string, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domains.User{}, "", constant.ErrUserNameNotFound

		}
		return domains.User{}, "", fmt.Errorf("account.GetUserByUsername: %w", err)
	}

	if user.IsActive == false {
		_, err := uc.GenerateAndSendOTP(
			ctx,
			SendOtpPayload{
				Email:  user.Email,
				Name:   user.Username,
				UserID: uint64(user.UserID),
			},
			true)
		if err != nil {
			log.Println(err)
			return domains.User{}, "", fmt.Errorf("uc.GenerateAndSendOTP: %w", err)
		}
		return domains.User{}, "", constant.ErrEmailIsNotVerified
	}

	chiperPass, err := hex.DecodeString(user.Password)
	if err != nil {
		log.Println(err)
		return domains.User{}, "", fmt.Errorf("hex.DecodeString: %w", err)
	}
	rawPass, err := uc.maker.DecryptMessage(
		[]byte(uc.env.Get("SECRET_ENCRYPTION_PASS")),
		chiperPass,
	)
	if err != nil {
		log.Println(err)
		if constant.ErrAesChipper.Error() == err.Error() {
			return domains.User{}, "", constant.ErrPasswordIsWrong
		}
		return domains.User{}, "", fmt.Errorf("uc.davinci.DecryptMessage: %w", err)
	}

	if string(rawPass) != payload.Password {
		return domains.User{}, "", constant.ErrPasswordIsWrong
	}

	issuedAt := jwt.NewNumericDate(uc.clock.Now(ctx))
	token, err := uc.auth.SignClaim(middleware.DefaultUserClaim{
		UserData: middleware.UserData{
			UserId:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   strconv.FormatUint(uint64(user.UserID), 10),
			IssuedAt: issuedAt,
		},
	})
	if err != nil {
		log.Println(err)
		return domains.User{}, "", fmt.Errorf("auth.SignClaim: %w", err)
	}

	return user, token, nil
}
