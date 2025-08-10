package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
	registerUser(user User) (RegisterUserResponse, error)
	getUserByUsername(username string) (User, error)
}

type service struct {
	repo      repository
	jwtSecret string
}

func newService(repo repository, jwtSecret string) service {
	return service{repo, jwtSecret}
}

func (s service) CreateUser(req RegisterUserRequest) (resp RegisterUserResponse, err error) {
	err = req.validate()
	if err != nil {
		return
	}

	hashedPasswordBytes, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		err = hashErr
		return
	}

	newUser := NewUser(req.Username, string(hashedPasswordBytes))
	createdUser, err := s.repo.registerUser(newUser)
	if err != nil {
		return
	}

	return createdUser, nil
}

func (s service) Login(req LoginRequest) (resp LoginResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}
	user, err := s.repo.getUserByUsername(req.Username)
	if err != nil {
		err = ErrInvalidCredentials
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		err = ErrInvalidCredentials
		return
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signErr := token.SignedString([]byte(s.jwtSecret))
	if signErr != nil {
		err = signErr
		return
	}
	resp.Token = tokenString
	return
}
