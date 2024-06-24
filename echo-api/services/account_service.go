package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kofkuiper/echo-api/config"
	"github.com/kofkuiper/echo-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type (
	accountService struct {
		cfg         config.Config
		accountRepo repositories.AccountRepository
	}
)

func NewAccountService(cfg config.Config, accountRepo repositories.AccountRepository) AccountService {
	return accountService{cfg: cfg, accountRepo: accountRepo}
}

// SignUp implements AccountService.
func (a accountService) SignUp(request SignUpRequest) (*Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	account, err := a.accountRepo.Create(repositories.Account{
		ID:       id.String(),
		Username: request.Username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:          account.ID,
		Username:    account.Username,
		CreatedAt:   account.CreatedAt,
		UpdatedAtAt: account.UpdatedAt,
	}, nil
}

// Login implements AccountService.
func (a accountService) Login(request LoginRequest) (string, error) {
	account, err := a.accountRepo.GetByUsername(request.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = validatePassword(account.Password, request.Password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := createJwtToken(request.Username, account.ID, a.cfg.JwtCfg)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Validate implements AccountService.
func (a accountService) Validate(token string) (*JwtClaims, error) {
	claims, err := ValidateJwtToken(token, a.cfg.JwtCfg.Secret)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func validatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func customJwtClaims(username, userId string, timeout time.Duration) JwtClaims {
	return JwtClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeout)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    fmt.Sprint(userId),
		},
	}
}

func createJwtToken(username, userId string, jwtCfg config.JwtConfig) (string, error) {
	claims := customJwtClaims(username, userId, jwtCfg.Timeout)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtCfg.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateJwtToken(tokenString, jwtSecretKey string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok {
		return claims, nil
	}
	return nil, errors.New("unexpected error")
}
