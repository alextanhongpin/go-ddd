package usecase

//go:generate string -type=ErrorCode
type ErrorCode int

const (
	ErrInvalidCredentials ErrorCode = iota
	ErrMissingAuthorizationHeader
	ErrUnauthorized
)
