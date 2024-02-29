package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/bowoBp/myDate/internal/constant"
	"github.com/bowoBp/myDate/internal/dto"
	"github.com/bowoBp/myDate/pkg/mapper"
	"log"
	"time"
)

type (
	Controller struct {
		uc     UseCaseInterface
		mapper mapper.MapperUtility
	}

	ControllerInterface interface {
		AddUser(
			ctx context.Context,
			payload RegisterPayload,
		) (*dto.Response, error)
		VerifyOtp(
			ctx context.Context,
			payload VerifyOtpPayload,
		) (*dto.Response, error)
		Login(
			ctx context.Context,
			payload LoginPayload,
		) (ResponseMetaLogin, error)
	}
)

func (ctrl Controller) AddUser(
	ctx context.Context,
	payload RegisterPayload,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.uc.AddUser(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Register is success",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) VerifyOtp(
	ctx context.Context,
	payload VerifyOtpPayload,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.uc.VerifyOtp(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !result {
		return nil, constant.ErrOtpInvalid
	}
	return dto.NewSuccessResponse(
		nil,
		"Email has been verified succesfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) Login(
	ctx context.Context,
	payload LoginPayload,
) (ResponseMetaLogin, error) {
	start := time.Now()

	user, token, err := ctrl.uc.Login(ctx, payload)

	loginData := LoginResponseData[LoginUserResponseData]{
		User: LoginUserResponseData{
			ID:       int(user.UserID),
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}

	if err != nil {
		log.Println(err)
		var (
			isVerified   = true
			isRegistered = true
			msgErr       = ""
		)
		switch {
		case errors.Is(err, constant.ErrUserNameNotFound):
			isRegistered = false
			isVerified = false
			msgErr = constant.ErrUserNameNotFound.Error()
			break
		case errors.Is(err, constant.ErrEmailIsNotVerified):
			isVerified = false
			msgErr = constant.ErrEmailIsNotVerified.Error()
			break
		default:
			break
		}
		return ResponseMetaLogin{
			Response: dto.Response{
				ResponseMeta: dto.ResponseMeta{
					Success:      false,
					Message:      msgErr,
					ResponseTime: fmt.Sprint(time.Since(start).Milliseconds(), ".ms"),
				},
				Data: loginData,
			},
			IsVerified:   isVerified,
			IsRegistered: isRegistered,
		}, err
	}

	return ResponseMetaLogin{
		Response: dto.Response{
			ResponseMeta: dto.ResponseMeta{
				Success:      true,
				MessageTitle: "Success",
				Message:      "berhasil login",
				ResponseTime: fmt.Sprint(time.Since(start).Milliseconds(), ".ms"),
			},
			Data: loginData,
		},
		IsVerified:   true,
		IsRegistered: true,
	}, err
}
