package repository

import (
	"database/sql"
	"errors"
	"my-hexagonal-api/internal/domain/user"
)

type PostgresUserRepository struct {
    db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
    return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(tx *sql.Tx, u *user.User) error {
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    return tx.QueryRow(query, u.Name, u.Email).Scan(&u.ID)
}

func (r *PostgresUserRepository) GetUserByID(id int64) (*user.User, error) {
    query := "SELECT id, name, email FROM users WHERE id = $1"
    row := r.db.QueryRow(query, id)

    user := &user.User{}
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}
