package domain

type CreateUserDto struct {
	Name string `validate:"required"`
}

type UpdateUserDto struct {
	Name string
}

type DeleteUserDto struct {
	ID string
}

type UserCommentsDto struct {
}
