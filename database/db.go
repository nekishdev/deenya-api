package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// var db *pgx.Conn
var db *sqlx.DB

func Init() {
	var err error
	t := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf(t, dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database:", dbName, "in", dbHost)
}
