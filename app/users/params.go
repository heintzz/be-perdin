package users

import "fmt"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func (req CreateUserRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("name field is required")
	}
	if req.Password == "" {
		return fmt.Errorf("email field is required")
	}
	return nil
}
