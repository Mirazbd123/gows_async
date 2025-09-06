package store

import (
	"context"
	"database/sql"
	"fmt"
	"gows-async/config"
	"time"
)

func NewPostgresDb(conf *config.Config) (*sql.DB, error) {
	dbUrl := conf.DatabaseUrl()
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db connections: %w", err)
	}
	return db, nil
}
