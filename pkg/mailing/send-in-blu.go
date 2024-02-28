package mailing

import (
	"context"
	"fmt"
	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type (
	SendInBlue struct {
		api *sendinblue.APIClient
	}

	SendInBlueInterface interface {
		GetAccount(ctx context.Context) (sendinblue.GetAccount, error)
		NativeSendEmail(ctx context.Context, payload NativeSendEmailPayload) error
		SendEmail(ctx context.Context, email sendinblue.SendSmtpEmail) (sendinblue.CreateSmtpEmail, error)
	}
)

func NewConfig() SendInBlueInterface {
	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: apis-key
	cfg.AddDefaultHeader("api-key", os.Getenv("SENDINBLUE_API_KEY"))
	//Configure API key authorization: partner-key
	cfg.AddDefaultHeader("partner-key", os.Getenv("SENDINBLUE_PARTNER_KEY"))
	sib := sendinblue.NewAPIClient(cfg)
	return &SendInBlue{
		api: sib,
	}
}

func (sib SendInBlue) NativeSendEmail(ctx context.Context, payload NativeSendEmailPayload) error {
	auth := smtp.PlainAuth("", payload.Username, payload.Password, payload.Host)
	messageBody := fmt.Sprintf(
		"From:  <%s>\n"+
			"To: <%s>\r\n"+
			"Subject: %s\r\n",
		payload.Username,
		payload.SendTo,
		payload.Subject,
	)
	messageBody += "MIME-version: 1.0;\r\n"
	messageBody += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	messageBody += payload.HtmlBody

	err := smtp.SendMail(
		payload.Host+":"+payload.Port,
		auth,
		payload.Username,
		[]string{payload.SendTo},
		[]byte(messageBody),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (sib SendInBlue) GetAccount(ctx context.Context) (sendinblue.GetAccount, error) {
	result, res, err := sib.api.AccountApi.GetAccount(ctx)
	if err != nil {
		log.Println(err)
		return sendinblue.GetAccount{}, err
	}
	if res.StatusCode != http.StatusOK {
		return sendinblue.GetAccount{}, fmt.Errorf("")
	}
	return result, nil

}

func (sib SendInBlue) SendEmail(
	ctx context.Context,
	email sendinblue.SendSmtpEmail,
) (sendinblue.CreateSmtpEmail, error) {
	result, _, err := sib.api.TransactionalEmailsApi.SendTransacEmail(ctx, email)
	if err != nil {
		log.Println(err)
		return sendinblue.CreateSmtpEmail{}, err
	}

	return result, nil
}
