package xerror

import "github.com/zeromicro/x/errors"

func XError(code int) error {
	return errors.New(code, MessageType[code])
}
