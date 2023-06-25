package main

import (
	"github.com/KPI-golang-5/Library/pkg/config"
	"github.com/KPI-golang-5/Library/pkg/controllers"
	. "github.com/KPI-golang-5/Library/pkg/repositories"
	"github.com/KPI-golang-5/Library/pkg/routes"
	"github.com/KPI-golang-5/Library/pkg/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	authorRepository := NewAuthorRepository(config.GetDB())
	bookRepository := NewBookRepository(config.GetDB())
	userRepository := NewUserRepository(config.GetDB())
	userFavBookRepository := NewUserFavBookRepository(config.GetDB())

	authorService := services.NewAuthorService(authorRepository)
	// add book, user and userFavBook SERVICES the same way

	authorController := controllers.RegisterAuthorController(authorService)
	// add book, user and userFavBook CONTROLLERS the same way

	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
