package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/p12s/library-rest-api/pkg/models"
	"github.com/p12s/library-rest-api/pkg/repository"
)

const (
	salt       = "8284kjalsdf282-4asfjae93sdf"
	tokenTTL   = 12 * time.Hour
	hmacSecret = "29dsjkadf*^(&le23#ls93s02a0d9"
)

// Authorization - signup/signin
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// AuthService - service
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService - constructor
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser - just creating user
func (s *AuthService) CreateUser(user models.User) (int, error) {
	passwordHash, err := s.generatePasswordHash(user.Password)
	if err != nil {
		return 0, fmt.Errorf("get user by creds: %w", err)
	}
	user.Password = passwordHash
	return s.repo.CreateUser(user)
}

// GenerateToken - token generation
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	passwordHash, err := s.generatePasswordHash(password)
	if err != nil {
		return "", fmt.Errorf("get user by creds: %w", err)
	}

	user, err := s.repo.GetUser(username, passwordHash)
	if err != nil {
		return "", fmt.Errorf("user creds wrong: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ //nolint
		Subject:   strconv.Itoa(user.Id),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(hmacSecret))
}

// ParseToken - getting authorized data from token

func (s *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return id, nil
}

// generatePasswordHash - hash generare from password
func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash := sha1.New() // #nosec
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("hash write: %w", err)
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
