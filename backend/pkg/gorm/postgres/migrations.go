package postgres

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(addr string) error {
	goose.SetBaseFS(embedMigrations)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return fmt.Errorf("failed to open postgres to migration, err %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect postgres to migration, err %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations, err %w", err)
	}
	return nil
}
