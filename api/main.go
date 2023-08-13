package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucas-simao/kubernetes/api/internal/domain/customers"
	"github.com/lucas-simao/kubernetes/api/internal/repository"
	api "github.com/lucas-simao/kubernetes/api/internal/server"
)

func main() {
	if _, ok := os.LookupEnv("IS_PROD"); !ok {
		err := godotenv.Load(".env")
		if err != nil {
			log.Panic("Error to load .env in the root directory")
		}
	}

	// Database
	repo := repository.New()

	// Apply migrations
	repo.ApplyMigration("./migrations")

	// Domains
	us := customers.New(repo)

	// Api
	a := api.New(us)
	api.Start(a)
}
