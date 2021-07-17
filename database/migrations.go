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
			created_at  INTEGER                           NOT NULL DEFAULT 0,
			updated_at  INTEGER                           NOT NULL DEFAULT 0
		)`
	if _, err := db.Exec(createTodoTable); err != nil {
		return errors.Wrap(err, "Cannot create todo table")
	}

	createDateTable := `
		CREATE TABLE date (
		  id    INTEGER                        PRIMEARY KEY,
		  date  TEXT CHECK(LENGTH(date) <= 10) NOT NULL DEFAULT (date('now')),
		  node  INTEGER                        NOT NULL,

		  FOREIGN KEY(date) REFERENCES todo (id) ON DELETE CASCADE
		  UNIQUE(date, node)
		)`
	if _, err := db.Exec(createDateTable); err != nil {
		return errors.Wrap(err, "Cannot create Date table")
	}

	createUpdateTrigger := `
		CREATE TRIGGER todo_updated_time UPDATE ON todo FOR EACH ROW
			BEGIN
				UPDATE todo SET updated_at = (strftime('%s', 'now')) WHERE id = old.id;
			END`
	if _, err := db.Exec(createUpdateTrigger); err != nil {
		return errors.Wrap(err, "Cannot create update trigger")
	}

	createEdgeTable := `
		CREATE TABLE edge (
		  id    INTEGER PRIMEARY KEY,
		  head  INTEGER NOT NULL,
		  tail  INTEGER NOT NULL,
		  root  INTEGER NOT NULL,

		  FOREIGN KEY(head) REFERENCES todo (id) ON DELETE CASCADE
		  FOREIGN KEY(tail) REFERENCES todo (id) ON DELETE CASCADE
		  FOREIGN KEY(root) REFERENCES todo (id) ON DELETE CASCADE

		  CHECK(head != tail)
		  UNIQUE(head, tail)
		  UNIQUE(root, tail)
		)`
	if _, err := db.Exec(createEdgeTable); err != nil {
		return errors.Wrap(err, "Cannot create edge table")
	}

	return nil
}
