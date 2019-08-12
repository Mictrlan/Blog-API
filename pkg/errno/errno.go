package errno

import "errors"

// error
var (
	ErrInconsistentPwd = errors.New("The entered password is inconsistent")
	ErrInvalidPwd      = errors.New("Wrong user name or password")
)
