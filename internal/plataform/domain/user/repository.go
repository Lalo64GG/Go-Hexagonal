package user

import (
    "database/sql"
    "errors"
)

type User struct {
    ID    int64
    Name  string
    Email string
}

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(tx *sql.Tx, user *User) error {
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    err := tx.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
    if err != nil {
        return err
    }
    return nil
}

func (r *UserRepository) GetUserByID(id int64) (*User, error) {
    query := "SELECT id, name, email FROM users WHERE id = $1"
    row := r.db.QueryRow(query, id)

    user := &User{}
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}
