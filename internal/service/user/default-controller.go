package user

import (
	"context"
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
	}
)

func (ctrl Controller) AddUser(
	ctx context.Context,
	payload RegisterPayload,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.uc.AddUser(ctx, payload)
	err = ctrl.mapper.EvaluateError("ctrl.Uc.Register", registerErrs, err)
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
	err = ctrl.mapper.EvaluateError("ctrl.Uc.VerifyOtp", verifyOtpErrs, err)
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
