package auth

type SignUpRequest struct {
	EmailID string `json:"emailId"`
}

type SignUpResponse struct {
	Token     string `json:"token"`
	EmailID   string `json:"emailId"`
	ExpiresAt string `json:"expiresAt"`
}

type VerifyAuthRequest struct {
	Token    string `json:"token"`
	EmailID  string `json:"emailId"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
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
	EmailID  string `json:"emailId"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string      `json:"accessToken"`
	Profile     UserProfile `json:"profile"`
}
