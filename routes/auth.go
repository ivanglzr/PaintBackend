package routes

import (
	"github.com/gorilla/mux"
	"github.com/ivanglzr/PaintBackend/config"
	"github.com/ivanglzr/PaintBackend/controllers"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc(config.Routes.Login, controllers.LogIn).Methods("POST")
	router.HandleFunc(config.Routes.Register, controllers.Register).Methods("POST")
	router.HandleFunc(config.Routes.DeleteUser, controllers.DeleteUser).Methods("DELETE")
}
