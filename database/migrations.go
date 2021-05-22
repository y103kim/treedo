package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

var migrationSequence = []func(*sql.DB) error{
	migrationV1,
}

func migrationV1(db *sql.DB) error {
	createTodoTable := `
		CREATE TABLE todo (
			id          INTEGER                           PRIMEARY KEY,
			title       TEXT CHECK(LENGTH(title) <= 100)  NOT NULL,
			status      TEXT CHECK(LENGTH(status) <= 20)  NOT NULL DEFAULT "Not Started",
			hidden      BOOLEAN                           NOT NULL DEFAULT 0,
			created_at  INTEGER                           NOT NULL DEFAULT (strftime('%s', 'now')),
			updated_at  INTEGER                           NOT NULL DEFAULT (strftime('%s', 'now'))
		)`
	if _, err := db.Exec(createTodoTable); err != nil {
		return errors.Wrap(err, "Cannot create table")
	}

	createUpdateTrigger := `
		CREATE TRIGGER todo_updated_time UPDATE ON todo FOR EACH ROW
			BEGIN
				UPDATE todo SET updated_at = (strftime('%s', 'now')) WHERE id = old.id;
			END`
	if _, err := db.Exec(createUpdateTrigger); err != nil {
		return errors.Wrap(err, "Cannot create update trigger")
	}

	return nil
}
