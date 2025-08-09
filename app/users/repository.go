package users

import "database/sql"

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return &userRepository{db: db}
}

func (r *userRepository) createUser(user User) (newUser CreateUserResponse, err error) {
	query := `INSERT INTO users (id, username, password, role, created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id, username, created_at`
	row := r.db.QueryRow(query, user.ID, user.Username, user.Password, user.Role, user.CreatedAt)
	err = row.Scan(&newUser.ID, &newUser.Username, &newUser.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (r *userRepository) updateUserRole(userID string, role string) (resp UpdateUserRoleResponse, err error) {
	query := `UPDATE users SET role = $1 WHERE id = $2 RETURNING id, username, role`
	row := r.db.QueryRow(query, role, userID)
	err = row.Scan(&resp.ID, &resp.Username, &resp.Role)
	if err != nil {
		return
	}
	return
}
