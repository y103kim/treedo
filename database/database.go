package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Database struct {
	db *sql.DB
}

func (database *Database) Open() error {
	db, err := sql.Open("sqlite3", "treedo.db")
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
