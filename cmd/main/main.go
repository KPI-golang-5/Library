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
	"os"
)

func main() {
	config.Connect()
	config.Migrate(config.GetDB())

	authorRepository := NewAuthorRepository(config.GetDB())
	bookRepository := NewBookRepository(config.GetDB())
	userRepository := NewUserRepository(config.GetDB())
	userFavBookRepository := NewUserFavBookRepository(config.GetDB())

	authorService := services.NewAuthorService(authorRepository)
	bookService := services.NewBookService(bookRepository)
	userService := services.NewUserService(userRepository)
	userFavBookService := services.NewUserFavBookService(userFavBookRepository)

	authorController := controllers.RegisterAuthorController(authorService)
	bookController := controllers.RegisterBookController(bookService)
	userController := controllers.RegisterUserController(userService)
	userFavController := controllers.RegisterUserFavBookController(userFavBookService)

	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r, authorController, bookController, userController, userFavController)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
