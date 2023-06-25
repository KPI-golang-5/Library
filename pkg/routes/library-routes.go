package routes

import (
	"github.com/KPI-golang-5/Library/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterLibraryRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/authors", authorController.GetAuthors).Methods("GET")
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/author/{authorId}", controllers.GetAuthorById).Methods("GET")
	router.HandleFunc("/author/{authorId}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{authorId}", controllers.DeleteAuthor).Methods("DELETE")

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/favBooks", controllers.GetAllFavBooks).Methods("GET")
	router.HandleFunc("/favBook", controllers.CreateFavBook).Methods("POST")
	router.HandleFunc("/favBook/{userId}", controllers.GetFavBooksByUserID).Methods("GET")
	router.HandleFunc("/favBook/{favBookId}", controllers.UpdateFavBook).Methods("PUT")
	router.HandleFunc("/favBook/{favBookId}", controllers.DeleteFavBook).Methods("DELETE")
}
