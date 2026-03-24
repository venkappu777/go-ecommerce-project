package database

import (
	"context" // Used for controlling requests
	"log"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn //This stores DB connection globally

func ConnectDB(){
	// pgx: PostgreSQL driver for Go just like mongoose for node
	conn,err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Dhoni@007@localhost:5432/ecommerce",
	)

	if err!=nil{
		log.Fatal("DB Connection Failed:", err)
	}

	DB = conn
	log.Println("Connected to PostgreSQL ✅")
}
