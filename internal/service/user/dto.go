package user

import "github.com/sendinblue/APIv3-go-library/v2/lib"

type (
	RegisterPayload struct {
		Name     string `json:"name" validate:"required"`
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
)
