package migrations

import (
	// "log"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up00001, Down00001)
}

func Up00001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		`
			CREATE TABLE users (
				id SERIAL,
				name VARCHAR(100),
				email VARCHAR(100) UNIQUE,
				password VARCHAR(255),
				age INT,
    			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
		`,
	)

	return err
}

func Down00001(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(
		ctx,
		"DROP TABLE users;",
	)
	return err
}
