package common

import "errors"

const (
	Success = 0

	InvalidArgument = 1000

	InternalServerError = 5000
)

var (
	ErrIDRequired     = errors.New("id is required")
	ErrDuplicatedKey  = errors.New("duplicated key")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
)
