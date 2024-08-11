package user

import (
    "my-hexagonal-api/internal/platform/database"
)

type UserService struct {
    repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(name, email string) (*User, error) {
    user := &User{Name: name, Email: email}

    err := database.WithTransaction(s.repo.db, func(tx *sql.Tx) error {
        return s.repo.CreateUser(tx, user)
    })

    if err != nil {
        return nil, err
    }

    return user, nil
}
