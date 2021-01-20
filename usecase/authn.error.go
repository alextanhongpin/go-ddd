package usecase

type AuthnErr int

//go:generate stringer -type=AuthnErr -trimprefix=AuthnErr
const (
	AuthnErrInvalidCredentials AuthnErr = iota
	AuthnErrMissingAuthorizationHeader
)
