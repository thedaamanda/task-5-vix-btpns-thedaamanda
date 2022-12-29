package request

type (
	CreateUserRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	UpdateUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
