package auth

import "database/sql"

type authRepository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return &authRepository{db: db}
}

func (r *authRepository) registerUser(user User) (newUser RegisterUserResponse, err error) {
	query := `INSERT INTO users (id, username, password, role, created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id, username, created_at`
	row := r.db.QueryRow(query, user.ID, user.Username, user.Password, user.Role, user.CreatedAt)
	err = row.Scan(&newUser.ID, &newUser.Username, &newUser.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (r *authRepository) getUserByUsername(username string) (user User, err error) {
	query := `SELECT id, username, password, role, created_at FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}
