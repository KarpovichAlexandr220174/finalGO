// finalGO/pkg/models/user.go
package models

import "finalGO/pkg/drivers"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	// Другие поля, если необходимо
}

// CreateUser сохраняет нового пользователя в базе данных
func CreateUser(user User) error {
	_, err := drivers.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username, user.Email, user.Password)
	return err
}

// AuthenticateUser аутентифицирует пользователя по имени пользователя и паролю
func AuthenticateUser(username, password string) (User, error) {
	var user User
	err := drivers.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1 AND password = $2",
		username, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	return user, err
}
