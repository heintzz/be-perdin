package users

type repository interface {
	createUser(user User) (CreateUserResponse, error)
}

type service struct {
	repo repository
}

func NewService(repo repository) service {
	return service{
		repo,
	}
}

func (s service) CreateUser(req CreateUserRequest) (user CreateUserResponse, err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	newUser := NewUser(req.Username, req.Password)
	createdUser, err := s.repo.createUser(newUser)
	if err != nil {
		return
	}

	return createdUser, nil
}
