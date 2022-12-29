package response

type PhotoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PhotosResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	User      PhotoUser `json:"User"`
}

type PhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhotoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    int    `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
