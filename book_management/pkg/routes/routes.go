package routes

import (
	"github.com/gorilla/mux"
	"book-management/pkg/controllers"
)

var RegisterBookMgmtRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}