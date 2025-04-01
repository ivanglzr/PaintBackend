package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanglzr/PaintBackend/db"
	"github.com/ivanglzr/PaintBackend/routes"
)

func main() {
	db.ConnectDatabase()

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
