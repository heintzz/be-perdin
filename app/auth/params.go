package auth

import "fmt"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func (req RegisterUserRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("username field is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password field is required")
	}
	return nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (req LoginRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("username field is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password field is required")
	}
	return nil
}

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")
