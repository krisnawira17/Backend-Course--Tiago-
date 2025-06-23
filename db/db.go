package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error){
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s",
        host, port, user, password, dbname,
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }

    return db, nil
}