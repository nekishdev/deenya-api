package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "PandoraTeam2021"
	DB_NAME     = "deenya"
	DB_HOST     = "173.212.240.109"
	DB_PORT     = "5432"
)

// var db *pgx.Conn
var db *sqlx.DB

func Init() {
	var err error
	t := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	connectionString := fmt.Sprintf(t, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database:", DB_NAME, "in", DB_HOST)
}

// func Init() {
// 	var err error
// 	var runtimeParams map[string]string
// 	runtimeParams = make(map[string]string)
// 	runtimeParams["application_name"] = "deenya"
// 	connConfig, err := pgx.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// connConfig.User = DB_USER
// 	// connConfig.Password = DB_PASSWORD
// 	// connConfig.Host = DB_HOST
// 	// connConfig.Port = DB_PORT
// 	// connConfig.Database = DB_NAME
// 	// connConfig.TLSConfig = nil
// 	// connConfig.RuntimeParams = runtimeParams

// 	db, err = pgx.ConnectConfig(context.Background(), connConfig)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
// 		os.Exit(1)
// 	}

// 	err = db.Ping(context.Background())

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to ping connection: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Connected to database:", DB_NAME, "in", DB_HOST)
// }
