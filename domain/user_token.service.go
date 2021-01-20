package domain

type UserTokenService interface {
	Sign(u User) (string, error)
	Verify(string) (User, error)
}
