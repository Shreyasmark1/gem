package auth

type SignUpRequest struct {
	EmailID string `json:"emailId" validate:"required,email"`
}

type SignUpResponse struct {
	Token     string `json:"token"`
	EmailID   string `json:"emailId"`
	ExpiresAt string `json:"expiresAt"`
}

type VerifyAuthRequest struct {
	Token    string `json:"token" validate:"required"`
	EmailID  string `json:"emailId" validate:"required,email"`
	FullName string `json:"fullName" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=8"`
}

type VerifyAuthResponse struct {
	AccessToken string      `json:"accessToken"`
	Profile     UserProfile `json:"profile"`
}

type UserProfile struct {
	UserID   string `json:"userId"`
	EmailID  string `json:"emailId"`
	FullName string `json:"fullName"`
}

type LoginRequest struct {
	EmailID  string `json:"emailId" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string      `json:"accessToken"`
	Profile     UserProfile `json:"profile"`
}
