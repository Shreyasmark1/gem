package auth

import (
	"context"
	"database/sql"
	"errors"
	"gem/internal/services/user"
	"gem/internal/store"
	"gem/internal/store/repository"
	"gem/pkg/response"
	"gem/pkg/utils"
	"log"
	"time"
)

type AuthService struct {
	userService *user.UserService
	store       *store.Store
	jwtSecret   string
}

func NewAuthService(store *store.Store, userService *user.UserService, jwtSecret string) *AuthService {
	return &AuthService{
		store:       store,
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (s *AuthService) InitiateAuth(ctx context.Context, emailID string) (*SignUpResponse, error) {

	_, err := s.store.Queries.GetUserByEmail(ctx, emailID)

	if err == nil {
		return nil, response.BadRequest("User with this email already exists. Please login instead.")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Printf("ERROR: GetUserByEmail failed for %s: %v", emailID, err)
		return nil, response.InternalServerError("Database error")
	}

	token, err := utils.GenerateVerificationToken()

	if err != nil {
		return nil, response.InternalServerError("Failed to generate token")
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	authTokenParam := repository.CreateAuthTokenParams{
		EmailID:   emailID,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	authToken, err := s.store.Queries.CreateAuthToken(ctx, authTokenParam)

	if err != nil {
		return nil, response.InternalServerError("Failed to save token")
	}

	return &SignUpResponse{
		Token:     authToken.Token,
		EmailID:   authToken.EmailID,
		ExpiresAt: expiresAt.Format(time.RFC3339),
	}, nil
}

func (s *AuthService) VerifyAuth(ctx context.Context, body VerifyAuthRequest) (*VerifyAuthResponse, error) {

	return store.WithTransaction(ctx, s.store, func(q *repository.Queries) (*VerifyAuthResponse, error) {

		authToken, err := q.FindAuthToken(ctx, repository.FindAuthTokenParams{
			Token:   body.Token,
			EmailID: body.EmailID,
		})

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, response.BadRequest("Invalid or expired token")
			}
			return nil, response.InternalServerError("Database error")
		}

		hashedPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			return nil, response.InternalServerError("Failed to hash password")
		}

		newUser, err := q.CreateUser(ctx, repository.CreateUserParams{
			FullName:     body.FullName,
			EmailID:      body.EmailID,
			PasswordHash: hashedPassword,
		})

		if err != nil {
			return nil, response.InternalServerError("Failed to create user")
		}

		err = q.MarkAuthTokenExpired(ctx, authToken.Token)

		if err != nil {
			return nil, err
		}

		jwtToken, err := utils.GenerateJWT(newUser.UserID.String(), s.jwtSecret)

		if err != nil {
			return nil, response.InternalServerError("Failed to generate token")
		}

		var profile = UserProfile{
			UserID:   newUser.UserID.String(),
			EmailID:  newUser.EmailID,
			FullName: newUser.FullName,
		}

		return &VerifyAuthResponse{
			AccessToken: jwtToken,
			Profile:     profile,
		}, nil
	})
}

func (s *AuthService) Login(ctx context.Context, body LoginRequest) (*LoginResponse, error) {

	user, err := s.store.Queries.GetUserByEmail(ctx, body.EmailID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.BadRequest("Invalid email or password")
		}
		return nil, response.InternalServerError("Database error")
	}

	if !utils.CheckPasswordHash(body.Password, user.PasswordHash) {
		return nil, response.BadRequest("Invalid email or password")
	}

	jwtToken, err := utils.GenerateJWT(user.UserID.String(), s.jwtSecret)

	if err != nil {
		return nil, response.InternalServerError("Failed to generate token")
	}

	var profile = UserProfile{
		UserID:   user.UserID.String(),
		EmailID:  user.EmailID,
		FullName: user.FullName,
	}

	return &LoginResponse{
		AccessToken: jwtToken,
		Profile:     profile,
	}, nil
}
