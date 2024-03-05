package main

import (
	"finalGO/pkg/models"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получение данных из формы входа
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Проверка введенных данных
		_, err := models.AuthenticateUser(username, password)
		if err != nil {
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

		// Создание нового пользователя
		user := models.User{Username: username, Password: password, Email: email}

		// Сохранение пользователя в базе данных
		err := models.CreateUser(user)
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
