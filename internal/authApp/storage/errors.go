package storage

import "errors"

var (
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrUserDoesNotExist    = errors.New("user does not exist")
	ErrServiceNotSupported = errors.New("service not supported")
	ErrTokenAlreadyExist   = errors.New("token/code already exist")
	ErrTokenDoesNotExist   = errors.New("token/code does not exist")
)
