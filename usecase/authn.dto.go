package usecase

type LoginDto struct {
	Email    string `json:"email" validate:"email,required" conform:"trim"`
	Password string `json:"password" validate:"min=6" conform:"trim"`
}

type RegisterDto struct {
	Email    string `json:"email" validate:"email,required" conform:"trim"`
	Password string `json:"password" validate:"min=6" conform:"trim"`
	Name     string `json:"name" validate:"required" conform:"trim"`
}
