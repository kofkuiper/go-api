package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AccountService interface {
		SignUp(SignUpRequest) (*Account, error)
		Login(LoginRequest) (string, error)
		Validate(string) (*JwtClaims, error)
	}

	Account struct {
		ID          string    `json:"id"`
		Username    string    `json:"username"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAtAt time.Time `json:"updated_at"`
	}

	SignUpRequest struct {
		Username string `json:"username" form:"username" validate:"required,max=32,min=6"`
		Password string `json:"password" form:"password" validate:"required,max=32,min=6"`
	}

	LoginRequest struct {
		Username string `json:"username" form:"username" validate:"required,max=32,min=6"`
		Password string `json:"password" form:"password" validate:"required,max=32,min=6"`
	}

	JwtClaims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
)
