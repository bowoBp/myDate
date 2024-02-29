package user

import "github.com/bowoBp/myDate/internal/constant"

var registerErrs = []error{
	constant.ErrEmailIsUsed,
	constant.ErrUserNameIsUsed,
	constant.ErrBranchNotFound,
	constant.ErrUserNameFormat,
	constant.ErrPhoneIsUsed,
	constant.ErrEmailIsUsedGeneric,
	constant.ErrPhoneIsUsedGeneric,
}

var verifyOtpErrs = []error{
	constant.ErrOtpInvalid,
	constant.ErrUserNameNotFound,
	constant.ErrNewOtpRequired,
}

var loginErrs = []error{
	constant.ErrUserNameNotFound,
	constant.ErrPasswordIsWrong,
	constant.ErrEmailIsNotVerified,
}
