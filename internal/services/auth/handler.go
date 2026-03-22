package auth

import (
	middleware "gem/internal/middlewares"
	"gem/pkg/response"
	"net/http"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) error {

	if body, err := middleware.DecodeAndValidate[SignUpRequest](r); err != nil {
		return response.BadRequest("Invalid request body")
	} else {
		res, err := h.authService.InitiateAuth(r.Context(), body.EmailID)

		if err != nil {
			return err
		}

		return response.Send(w, res)
	}

}

func (h *AuthHandler) VerifyAuth(w http.ResponseWriter, r *http.Request) error {
	if body, err := middleware.DecodeAndValidate[VerifyAuthRequest](r); err != nil {
		return response.BadRequest("Invalid request body")
	} else {
		res, err := h.authService.VerifyAuth(r.Context(), body)

		if err != nil {
			return err
		}

		return response.Send(w, res)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	if body, err := middleware.DecodeAndValidate[LoginRequest](r); err != nil {
		return response.BadRequest("Invalid request body")
	} else {
		res, err := h.authService.Login(r.Context(), body)

		if err != nil {
			return err
		}

		return response.Send(w, res)
	}
}
