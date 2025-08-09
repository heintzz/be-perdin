package users

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
	updateUserRole(userID string, role string) (UpdateUserRoleResponse, error)
	getUserByID(userID string) (GetUserProfileResponse, error)
	getUserByUsername(username string) (User, error)
}

type service struct {
	repo      repository
	jwtSecret string
}

func NewService(repo repository, jwtSecret string) service {
	return service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s service) UpdateUserRole(req UpdateUserRoleRequest) (resp UpdateUserRoleResponse, err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	return s.repo.updateUserRole(req.UserID, req.Role)
}

func (s service) GetUserProfile(req GetUserProfileRequest) (resp GetUserProfileResponse, err error) {
	err = req.Validate()
	if err != nil {
		return
	}
	return s.repo.getUserByID(req.UserID)
}

func (s service) Login(req LoginRequest) (resp LoginResponse, err error) {
	if err = req.Validate(); err != nil {
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
