package usecase

type ErrorCode int

//go:generate stringer -type=ErrorCode
const (
	ErrInvalidCredentials ErrorCode = iota
	ErrMissingAuthorizationHeader
	ErrUnauthorized
)
