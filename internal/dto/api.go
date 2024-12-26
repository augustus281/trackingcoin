package dto

type RegisterRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetDetailCMCRequest struct {
	Slug string `json:"slug"`
}
