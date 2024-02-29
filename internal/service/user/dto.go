package user

import (
	"github.com/bowoBp/myDate/internal/dto"
	"github.com/sendinblue/APIv3-go-library/v2/lib"
)

type (
	RegisterPayload struct {
		UserName string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"required"`
	}

	RegisterResponseData struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		IsActive bool   `json:"isActive"`
	}

	RegisterResult struct {
		User  RegisterResponseData `json:"user"`
		Email lib.CreateSmtpEmail  `json:"emailResult"`
	}

	SendOtpPayload struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Content string `json:"content"`
		Subject string `json:"subject"`
		UserID  uint64 `json:"userId" validate:"numeric"`
	}

	SendOtpResult struct {
		Otp         int                 `json:"otp"`
		EmailResult lib.CreateSmtpEmail `json:"emailResult"`
	}

	VerifyOtpPayload struct {
		Otp    string `json:"otpCode" validate:"numeric"`
		UserID uint64 `json:"userId" validate:"numeric"`
	}

	LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password" validate:"required"`
	}
	ResponseMetaLogin struct {
		dto.Response
		IsVerified   bool `json:"isVerified"`
		IsRegistered bool `json:"isRegistered"`
	}

	LoginResponseData[Data any] struct {
		User  Data   `json:"user"`
		Token string `json:"token"`
	}

	LoginUserResponseData struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
