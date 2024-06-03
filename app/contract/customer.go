package contract

import "errors"

var (
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrWrongPassword        = errors.New("wrong password")
)

type CustomerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
