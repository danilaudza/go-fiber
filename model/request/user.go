package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address"`
	Phonee  string `json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phonee  string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"name" validate:"required"`
}
