package main

import (
	"log"

	"github.com/krisnawira17/go-backend-learn/cmd/api"
	"github.com/krisnawira17/go-backend-learn/cmd/config"
	"github.com/krisnawira17/go-backend-learn/db"
)

func main() {
	cfg := config.LoadConfig()
	conn:= db.NewPostgresConnection(cfg)
	defer conn.Close()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}