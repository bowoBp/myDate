package global_usage

type EmailBodyVerifyOTPPayload struct {
	Name       string   `json:"name"`
	OTPs       []string `json:"otps"`
	VerifyPage string   `json:"verifyPage"`
}
