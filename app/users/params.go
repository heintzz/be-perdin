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
		return fmt.Errorf("username field is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password field is required")
	}
	return nil
}

type UpdateUserRoleRequest struct {
	UserID string `json:"-"`
	Role   string `json:"role"`
}

type UpdateUserRoleResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (req UpdateUserRoleRequest) Validate() error {
	if req.UserID == "" {
		return fmt.Errorf("userId param is required")
	}
	if req.Role == "" {
		return fmt.Errorf("role field is required")
	}
	if req.Role != "PEGAWAI" && req.Role != "SDM" {
		return fmt.Errorf("invalid role")
	}
	return nil
}

type GetUserProfileRequest struct {
	UserID string `json:"-"`
}

type GetUserProfileResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

func (req GetUserProfileRequest) Validate() error {
	if req.UserID == "" {
		return fmt.Errorf("user id is required")
	}
	return nil
}
