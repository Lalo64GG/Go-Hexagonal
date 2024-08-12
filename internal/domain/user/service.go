package user

import (
    "database/sql"
    "my-hexagonal-api/internal/platform/database"
)

type UserService struct {
    repo UserRepository
    db   *sql.DB
}

func NewUserService(repo UserRepository, db *sql.DB) *UserService {
    return &UserService{repo: repo, db: db}
}

func (s *UserService) RegisterUser(name, email string) (*User, error) {
    user := &User{Name: name, Email: email}

    err := database.WithTransaction(s.db, func(tx *sql.Tx) error {
        return s.repo.CreateUser(tx, user)
    })

    if err != nil {
        return nil, err
    }

    return user, nil
}
