package common

import "errors"

const (
	Success = 0

	InvalidArgument = 1000

	InternalServerError = 5000
)

var (
	ErrMissingField   = errors.New("missing field")
	ErrDuplicatedKey  = errors.New("duplicated key")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
)
