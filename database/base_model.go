package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type BaseModel interface {
	SetId(id int64)
	GetTableName() string
	GetFieldNames() string
	GetValueList() string
	GetUpdateList(fields []string) string
	Deserialize(db_output string) error
}

func (db *Database) Insert(obj BaseModel) error {
	table := obj.GetTableName()
	fields := obj.GetFieldNames()
	values := obj.GetValueList()
	return db.Tx(func(tx *sql.Tx) error {
		cmd := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fields, values)
		r, err := tx.Exec(cmd)
		if err != nil {
			return errors.Wrapf(err, "Error while exec SQL\n%s\n", cmd)
		}
		id, _ := r.LastInsertId()
		obj.SetId(id)
		return nil
	})
}
