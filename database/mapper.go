package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Mapper interface {
	SetId(id int64)
	GetTableName() string
	GetFieldNames() string
	GetValueList() string
	GetPkFieldName() string
	GetUpdateList(fields []string) string
}

func (db *Database) Insert(obj Mapper) error {
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

func (db *Database) Get(obj Mapper, id int64) error {
	table := obj.GetTableName()
	pk_name := obj.GetPkFieldName()
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", table, pk_name, id)
	return errors.Wrapf(db.db.Get(obj, cmd), "Error whlie select\n%s\n", cmd)
}
