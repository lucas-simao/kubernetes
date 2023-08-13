package repository

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New() Repository {
	user, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		log.Panic("Error to get DATABASE_USER")
	}
	password, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		log.Panic("Error to get DATABASE_PASSWORD")
	}
	name, ok := os.LookupEnv("DATABASE_DB_NAME")
	if !ok {
		log.Panic("Error to get DATABASE_DB_NAME")
	}
	host, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		log.Panic("Error to get DATABASE_HOST")
	}
	port, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		log.Panic("Error to get DATABASE_PORT")
	}

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)

	newDb, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Panic(err)
	}

	return &repository{
		db: newDb,
	}
}

func (r *repository) ApplyMigration(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".up.sql") {
			file, err := os.ReadFile(fmt.Sprintf("%s/%s", path, f.Name()))
			if err != nil {
				log.Panic(err)
			}

			_, err = r.db.Exec(string(file))
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
