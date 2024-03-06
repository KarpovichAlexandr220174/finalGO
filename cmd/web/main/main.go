// final.go/cmd/web/main/main.go
package main

import (
	"finalGO/pkg/drivers"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	drivers.InitDB("user=postgres dbname=finalGo password=0000 sslmode=disable\n")

	// Инициализация маршрутов с использованием mux
	mux := http.NewServeMux()

	//----главная страница
	mux.HandleFunc("/", home)
	//----логин и регистрация
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	//----страница после успешного входа
	mux.HandleFunc("/application", application)
	//----номер и имя
	mux.HandleFunc("/city/", searchPageHandler)
	mux.HandleFunc("/search/", searchHandler)
	//----больницы
	mux.HandleFunc("/city/hospitals/", searchHospitalsPageHandler)
	mux.HandleFunc("/search/hospitals/", searchHospitalsHandler)

	mux.HandleFunc("/city/schools/", searchSchoolsPageHandler)
	mux.HandleFunc("/search/schools/", searchSchoolsHandler)
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
