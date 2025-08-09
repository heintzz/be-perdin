package users

import "database/sql"

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return &userRepository{db: db}
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

func (r *userRepository) getUserByID(userID string) (resp GetUserProfileResponse, err error) {
	query := `SELECT id, username, role, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, userID)
	err = row.Scan(&resp.ID, &resp.Username, &resp.Role, &resp.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (r *userRepository) getUserByUsername(username string) (user User, err error) {
	query := `SELECT id, username, password, role, created_at FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}
