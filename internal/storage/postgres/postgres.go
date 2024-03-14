package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	cfg "github.com/paniccaaa/adamenblog-api/internal/config"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgres(cfg *cfg.ConfigDatabase) (*PostgresStore, error) {
	const op = "storage.postgres.NewPostgres"
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PostgresStore{db: db}, nil
}
