package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Database struct {
	db      *sql.DB
	Version int
}

func (database *Database) Open(fileName string) error {
	db, err := sql.Open("sqlite3", fileName)
	database.db = db
	return errors.Wrap(err, "Fail to open db")
}

type TxCb func(*sql.Tx) error

func (database *Database) Tx(cb TxCb) error {
	ctx := context.TODO()
	tx, err := database.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return errors.Wrap(err, "Cannot begin transction")
	}
	if err := cb(tx); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Error while transaction")
	}
	return tx.Commit()
}

func (database *Database) GetVersion() error {
	row := database.db.QueryRow("PRAGMA user_version")
	err := row.Scan(&database.Version)
	return errors.Wrap(err, "Cannot read PRAGMA user_version")
}

func (database *Database) SetVersion(version int) error {
	database.Version = version
	query := fmt.Sprintf("PRAGMA user_version = %d", version)
	_, err := database.db.Exec(query)
	return errors.Wrapf(err, "Fail to set PRAGMA user_version as %d", version)
}
