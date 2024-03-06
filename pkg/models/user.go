// finalGO/pkg/models/user.go
package models

import (
	"finalGO/pkg/drivers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	// Другие поля, если необходимо
}

// CreateUser сохраняет нового пользователя в базе данных
func CreateUser(user User) error {
	// Хеширование пароля перед сохранением в базе данных
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = drivers.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username, user.Email, hashedPassword)
	return err
}

func GetHashedPassword(username string) (string, error) {
	var hashedPassword string
	err := drivers.DB.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	return hashedPassword, err
}
