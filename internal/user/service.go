package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GetByEmail(email string) (User, error)
	VerifyUserPassword(user User, loginRequest LoginRequest) error
	GenerateToken(user User, secret string, duration time.Duration) (string, error)
	Register(registerRequest RegisterRequest) (User, error)
	GetUserToken(r *http.Request) (string, error)
	GetUserByToken(token string) (User, error)
}

type service struct {
	repo Repository
}

var UserService Service

func NewService(repo Repository) Service {
	UserService = &service{
		repo: repo,
	}
	return UserService
}

func (s *service) GetByEmail(email string) (User, error) {
	// get from the database
	user, err := s.repo.GetByEmail(email)
	return user, err
}

func (s *service) VerifyUserPassword(user User, loginRequest LoginRequest) error {
	// TODO use hashing for password
	if user.Password != loginRequest.Password {
		return errors.New("password not match")
	}
	return nil
}

func (s *service) GenerateToken(user User, secret string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(time.Hour * duration).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) Register(registerRequest RegisterRequest) (User, error) {
	user := User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password, // TODO hash
	}
	user, err := s.repo.Save(user)
	return user, err
}

func (s *service) GetUserToken(r *http.Request) (string, error) {
	token := r.Header.Get("authorization")
	if token == "" {
		return "", errors.New("missing auth token")
	}
	return token, nil
}
func (s *service) GetUserByToken(token string) (User, error) {
	// TODO add decode jwt token
	// get user by id
	var userID uint = 1
	user, err := s.repo.GetByID(userID)
	return user, err
}
