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
