package common

import "errors"

const (
	Success = 0

	InvalidArgument = 1000
)

var (
	ErrIDRequired = errors.New("id is required")
)
