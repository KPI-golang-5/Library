package routes

import (
	. "github.com/KPI-golang-5/Library/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterLibraryRoutes = func(router *mux.Router, authorController *AuthorController, bookController *BookController,
	userController *UserController, userFavBookController *UserFavBookController) {
	router.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/book", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", bookController.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bookController.DeleteBook).Methods("DELETE")

	router.HandleFunc("/authors", authorController.GetAuthors).Methods("GET")
	router.HandleFunc("/author", authorController.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/{authorId}", authorController.GetAuthorById).Methods("GET")
	router.HandleFunc("/author/{authorId}", authorController.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{authorId}", authorController.DeleteAuthor).Methods("DELETE")

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", userController.DeleteUser).Methods("DELETE")

	router.HandleFunc("/favBooks", userFavBookController.GetAllFavBooks).Methods("GET")
	router.HandleFunc("/favBook", userFavBookController.CreateFavBook).Methods("POST")
	router.HandleFunc("/favBook/{userId}", userFavBookController.GetFavBooksByUserId).Methods("GET")
	router.HandleFunc("/favBook/{favBookId}", userFavBookController.UpdateFavBook).Methods("PUT")
	router.HandleFunc("/favBook/{favBookId}", userFavBookController.DeleteFavBook).Methods("DELETE")
}
