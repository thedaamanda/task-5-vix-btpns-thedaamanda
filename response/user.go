package response

type UserResponse struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	UpdatedAt string `json:"updated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
