package routes

import (
	"github.com/gorilla/mux"
	"github.com/ivanglzr/PaintBackend/controllers"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/log-in", controllers.LogIn).Methods("POST")
	router.HandleFunc("/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/auth", controllers.DeleteUser).Methods("DELETE")
}
