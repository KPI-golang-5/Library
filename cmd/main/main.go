package main

import (
	"github.com/KPI-golang-5/Library/pkg/controllers"
	"github.com/KPI-golang-5/Library/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLibraryRoutes(r)
	r.Use(controllers.JwtAuthentication)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
