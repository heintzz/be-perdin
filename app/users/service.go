package users

import (
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
	createUser(user User) (CreateUserResponse, error)
	updateUserRole(userID string, role string) (UpdateUserRoleResponse, error)
	getUserByID(userID string) (GetUserProfileResponse, error)
}

type service struct {
	repo repository
}

func NewService(repo repository) service {
	return service{
		repo,
	}
}

func (s service) CreateUser(req CreateUserRequest) (resp CreateUserResponse, err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	hashedPasswordBytes, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		err = hashErr
		return
	}

	newUser := NewUser(req.Username, string(hashedPasswordBytes))
	createdUser, err := s.repo.createUser(newUser)
	if err != nil {
		return
	}

	return createdUser, nil
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
