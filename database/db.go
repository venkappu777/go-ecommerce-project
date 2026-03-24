package database

import (
	"context" // Used for controlling requests
	"log"
	"os"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn //This stores DB connection globally

func ConnectDB(){
	dbURL := os.Getenv("DB_URL")
	// pgx: PostgreSQL driver for Go just like mongoose for node
	conn,err := pgx.Connect(
		context.Background(),
		dbURL,
	)

	if err!=nil{
		log.Fatal("DB Connection Failed:", err)
	}

	DB = conn
	log.Println("Connected to PostgreSQL ✅")
}
