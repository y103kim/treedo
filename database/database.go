package database

import (
	"fmt"

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
	if err == nil {
		database := &Database{client}
		return database
	} else {
		panic("Cannot open or create database")
	}
}

func (d *Database) Close() error {
	return errors.Wrap(d.client.Close(), "Cannot close database")
}
