package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"money-management-be/database"
	"money-management-be/database/migrations"
	"money-management-be/routes"
)

func main() {
	// inisialisasi environtment variabel
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
	appPort := os.Getenv("APP_PORT")
	appServer := os.Getenv("APP_SERVER")

	// konfigurasi database
	database.Connect()
	migrations.Migrate()

	// inisialisasi routing
	mux := routes.InitRoutes(database.DB)

	// jalankan REST API
	log.Println("Server running at http://" + appServer + appPort)
	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatal(err)
	}
}
