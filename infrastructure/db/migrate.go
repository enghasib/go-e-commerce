package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate() {

	dbURL := "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable"

	m, err := migrate.New("file://db/migrations", dbURL)

	if err != nil {
		fmt.Println("Migration error:", err)
		log.Fatal("Migration failed!", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully!")
}
