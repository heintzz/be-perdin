package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
}

func NewUser(username, password string) User {
	currentTime := time.Now()
	return User{
		ID:        uuid.NewString(),
		Username:  username,
		Password:  password,
		Role:      "PEGAWAI",
		CreatedAt: currentTime.Format(time.RFC3339),
	}
}
