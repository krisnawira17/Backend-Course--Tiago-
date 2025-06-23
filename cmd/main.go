package main

import (
	"log"

	"github.com/krisnawira17/go-backend-learn/cmd/api"
	"github.com/krisnawira17/go-backend-learn/db"
)

func main() {
	_, err := db.NewPostgresStorage()
	if err != nil{
		log.Fatalf("Fail to conenct to DB: %v", err)
	}

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}