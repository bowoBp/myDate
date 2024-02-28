package constant

import "fmt"

var (
	ErrUserNameIsUsed      = fmt.Errorf("your email already use")
	ErrResendOTPNotAllowed = fmt.Errorf("your email has benn verified")
	ErrEmailIsUsed         = fmt.Errorf("Email sudah digunakan oleh owner lainnya")
	ErrBranchNotFound      = fmt.Errorf("Data outlet tidak ditemukan")
	ErrUserNameFormat      = fmt.Errorf("Username tidak dapat berisi spasi atau/dan dimulai dengan angka")
	ErrPhoneIsUsed         = fmt.Errorf("Nomor telepon telah digunakan oleh pengguna lainnya")
	ErrPhoneIsUsedGeneric  = fmt.Errorf("Phone {0} sudah digunakan oleh akun lain")
	ErrEmailIsUsedGeneric  = fmt.Errorf("Email {0} sudah digunakan oleh akun lain")
	ErrUserNameNotFound    = fmt.Errorf("Username atau email anda tidak ditemukan")
	ErrNewOtpRequired      = fmt.Errorf("Kode OTP anda telah kedaluwarsa")
	ErrOtpInvalid          = fmt.Errorf("Kode OTP yang anda masukan salah")
)
