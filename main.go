package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ivanglzr/PaintBackend/db"
	"github.com/ivanglzr/PaintBackend/middlewares"
	"github.com/ivanglzr/PaintBackend/routes"
)

func main() {
	db.ConnectDatabase()

	router := mux.NewRouter()

	router.Use(middlewares.AuthMiddleware)

	routes.AuthRoutes(router)

	port := os.Getenv("PORT")
	portStr := fmt.Sprintf(":%v", port)

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}
