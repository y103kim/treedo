package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Database struct {
	db       *sql.DB
	fileName string
	version  int
}

func (db *Database) Open(fileName string) error {
	sqldb, err := sql.Open("sqlite3", fileName)
	db.db, db.fileName = sqldb, fileName
	return errors.Wrap(err, "Fail to open db")
}

type TxCb func(*sql.Tx) error

func (db *Database) Tx(cb TxCb) error {
	ctx := context.TODO()
	tx, err := db.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return errors.Wrap(err, "Cannot begin transction")
	}
	if err := cb(tx); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Error while transaction")
	}
	return tx.Commit()
}

func (db *Database) getVersion() error {
	row := db.db.QueryRow("PRAGMA user_version")
	err := row.Scan(&db.version)
	return errors.Wrap(err, "Cannot read PRAGMA user_version")
}

func (db *Database) setVersion(version int) error {
	db.version = version
	query := fmt.Sprintf("PRAGMA user_version = %d", version)
	_, err := db.db.Exec(query)
	return errors.Wrapf(err, "Fail to set PRAGMA user_version as %d", version+1)
}

func (db *Database) migrateTo(version int) error {
	migration := migrationSequence[version]
	if err := migration(db.db); err != nil {
		return errors.Wrapf(err, "Fail to migrate to version %d", version+1)
	}
	return db.setVersion(version + 1)
}

func (db *Database) Migrate() error {
	currentVersion := db.version
	targetVersion := len(migrationSequence)
	var err error = nil
	for v := currentVersion; v < targetVersion && err == nil; v++ {
		err = db.migrateTo(v)
	}
	return err
}

func (db *Database) Close() {
	db.db.Close()
}
