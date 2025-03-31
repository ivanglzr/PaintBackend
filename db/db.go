package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	//TODO: use .env for dsn
	dsn := "host=localhost user=postgres password=postgres dbname=paint_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo hacer ping a la base de datos:", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
	DB = db
}
