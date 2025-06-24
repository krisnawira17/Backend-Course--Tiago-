package migrate

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/krisnawira17/go-backend-learn/cmd/config"
)

func RunMigrations(cfg config.Config){
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,)
	fmt.Printf(dbURL)
	m, err := migrate.New(
		"file://cmd/migrate/migrations", 
		dbURL,
	)
	if err != nil {
		log.Fatalf("migration init failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}