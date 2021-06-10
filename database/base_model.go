package database

import (
	"database/sql"

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

func (db *Database) Insert(obj *BaseModel) error {
	table := (*obj).GetTableName()
	fields := (*obj).GetFieldNames()
	values := (*obj).GetValueList()
	return db.Tx(func(tx *sql.Tx) error {
		r, err := tx.Exec("INSERT INTO ? (?) VALUES (?)", table, fields, values)
		if err != nil {
			return errors.Wrapf(err, "Cannot insert object %v", obj)
		}
		id, _ := r.LastInsertId()
		(*obj).SetId(id)
		return nil
	})
}
