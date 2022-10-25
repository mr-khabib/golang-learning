package routes

import (
	"github.com/gorilla/mux"
	"github.com/mr-khabib/golang-learning/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
}
