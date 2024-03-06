package main

import (
	"finalGO/pkg/drivers"
	"finalGO/pkg/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получение данных из формы входа
		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, err := models.GetHashedPassword(username)
		if err != nil {
			fmt.Printf("Error getting hashed password: %v\n", err)
			http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
			return
		}

		fmt.Printf("Hashed Password from DB: %v\n", hashedPassword)

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				// Неправильный пароль
				fmt.Printf("Invalid password: %v\n", err)
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				return
			}
			// Другие ошибки при сравнении пароля
			fmt.Printf("Error comparing passwords: %v\n", err)
			http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
			return
		}

		// Редирект на защищенную страницу
		http.Redirect(w, r, "/application", http.StatusSeeOther)
		return
	}

	// Отображение страницы входа
	tmpl.ExecuteTemplate(w, "login.tmpl", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получение данных из формы регистрации
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		// Проверка уникальности логина
		var exists bool
		err := drivers.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
		if err != nil {
			http.Error(w, "Error checking for existing username", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		// Проверка уникальности почты
		err = drivers.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
		if err != nil {
			http.Error(w, "Error checking for existing email", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}

		// Создание нового пользователя
		user := models.User{Username: username, Password: password, Email: email}

		// Сохранение пользователя в базе данных
		err = models.CreateUser(user)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Редирект на главную страницу или другую страницу
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Отображение страницы регистрации
	tmpl.ExecuteTemplate(w, "register.tmpl", nil)
}
