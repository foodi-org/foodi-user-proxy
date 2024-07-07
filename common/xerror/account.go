package xerror

import "github.com/zeromicro/x/errors"

var (
	InvalidRegisterType = errors.New(100000, "register type invalid")

	InvalidPhone = errors.New(100001, "phone invalid")
	MissPhone    = errors.New(100, "miss phone number")
	MissCode     = errors.New(101, "miss code number")
	MissPassword = errors.New(102, "miss password number")

	RegisterFail = errors.New(103, "register fail")
)
