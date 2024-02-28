package global_usage

import (
	"bytes"
	"html/template"
	"log"
)

type (
	Usecase struct {
	}

	UsecaseInterface interface {
		GenerateEmailBodyVerifyOTP(
			payload EmailBodyVerifyOTPPayload,
		) (string, error)
	}
)

func NewUsecase() UsecaseInterface {
	return Usecase{}
}

func (uc Usecase) GenerateEmailBodyVerifyOTP(
	payload EmailBodyVerifyOTPPayload,
) (string, error) {
	htmlPath := "./_template/mailing/verification-email.html"
	tmpl := template.Must(template.ParseFiles(htmlPath))
	outWriter := bytes.Buffer{}

	err := tmpl.Execute(&outWriter, payload)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return outWriter.String(), nil
}
