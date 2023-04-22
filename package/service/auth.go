package service

import (
	"time"

	"github.com/Tasma-110110/mid1-prj"
	"github.com/Tasma-110110/mid1-prj/package/repository"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

const (
	signingKey = "sajsijffq6xbj$%snr343dfdh"
	tokenTTL   = 2 * time.Hour
)

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) CreateUser(user mid1.User) (int, error) {
	return s.repo.CreateUser(user)
}
