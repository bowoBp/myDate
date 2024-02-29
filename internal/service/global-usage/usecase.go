package global_usage

import (
	"bytes"
	"github.com/bowoBp/myDate/pkg/environment"
	"github.com/bowoBp/myDate/pkg/maker"
	"html/template"
	"log"
)

type (
	Usecase struct {
		env   environment.Environment
		maker maker.Generator
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
