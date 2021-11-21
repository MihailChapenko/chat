package postgres

import (
	"github.com/MihailChapenko/chat/db"
	"github.com/MihailChapenko/chat/internal/models"
	"github.com/jmoiron/sqlx"
)

//UserRepository user repository interface
type UserRepository interface {
	FindByUsername(username string) (models.User, error)
	Create(models.User) (models.User, error)
}

//userRepository describe user repository struct
type userRepository struct {
	conn *sqlx.DB
}

//NewUserRepository create new user repository instance
func NewUserRepository() UserRepository {
	c := db.GetDB()

	return &userRepository{
		conn: c,
	}
}

//FindByUsername find user by username
func (u userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := u.conn.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

//Create new user
func (u userRepository) Create(user models.User) (models.User, error) {
	rows, err := u.conn.NamedQuery(u.conn.Rebind(
		`INSERT INTO users (username,password) VALUES (:username,:password) RETURNING id`), user)
	if err != nil {
		return models.User{}, err
	}
	if rows.Next() {
		rows.Scan(&user.ID)
	}

	return user, nil
}
