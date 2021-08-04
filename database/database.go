package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/y103kim/treedo/ent"
)

type Database struct {
	client *ent.Client
}

func CreateDatabase(filename string) *Database {
	cmd := fmt.Sprintf("file:%s?cache=shared&_fk=1", filename)
	client, err := ent.Open("sqlite3", cmd)
	if err != nil {
		log.Fatal("Cannot open or create database")
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	database := &Database{client}
	return database
}

func (d *Database) Close() error {
	return errors.Wrap(d.client.Close(), "Cannot close database")
}

type TxCb func(ctx context.Context, tx *ent.Tx) error

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func (d *Database) Tx(fn TxCb) error {
	ctx := context.Background()
	tx, err := d.client.Tx(ctx)
	if err != nil {
		return err
	}
	if err := fn(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return errors.Wrap(tx.Commit(), "committing transaction:")
}
