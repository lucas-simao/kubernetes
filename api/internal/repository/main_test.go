package repository

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/lucas-simao/kubernetes/api/configs"
)

var (
	repo Repository
	DB   *sqlx.DB
)

func TestMain(m *testing.M) {
	port := "5532"
	newContainer := configs.ContainerRun(port)
	// newContainer.RunMigrations("../../scripts/migrations")

	os.Setenv("DATABASE_USER", "postgres")
	os.Setenv("DATABASE_PASSWORD", "postgres")
	os.Setenv("DATABASE_DB_NAME", "postgres")
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", port)

	repo = New()
	DB = newContainer.DB

	repo.ApplyMigration("../../scripts/migrations")

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}
