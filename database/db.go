package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// inisialisasi environtment variabel untuk database
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	// inisialisasi database postgre
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password,
	)

	// buka database postgre
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// cek ada atau tidaknya database
	var exists bool
	query := `SELECT EXISTS(SELECT datname FROM pg_database WHERE datname = $1);`
	err = DB.QueryRow(query, dbname).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking database existance: ", err)
	}

	// kondisi jika database belum dibuat
	if !exists {
		_, err = DB.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
		if err != nil {
			log.Fatal("Error Creating database: ", err)
		}
		fmt.Println("Database", dbname, "created successfully!")
	}

	// buka database
	psqlDefault := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", psqlDefault)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connected successfully!")
}
