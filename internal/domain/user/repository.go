package user
import(
    "database/sql"
)

type User struct {
    ID    int64
    Name  string
    Email string
}

type UserRepository interface {
    CreateUser(tx *sql.Tx, user *User) error
    GetUserByID(id int64) (*User, error)
}
