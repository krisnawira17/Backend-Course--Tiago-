package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/krisnawira17/go-backend-learn/cmd/config"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg config.Config) *sql.DB{
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,)
    db, err := sql.Open("postgres", dsn)
    if err != nil{
        log.Fatal("Failed to open DB connection: ", err)
    }
    err = db.Ping()
    if err != nil{
        log.Fatal("Failed to ping DB: ", err)
    }
    fmt.Println("Conneted to PostgreSQL")
    return db
}