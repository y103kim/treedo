package database

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Mapper interface {
	SetId(id int64)
	TableName() string
	IdFieldName() string
}

func (db *Database) Insert(obj Mapper) error {
	table := obj.TableName()
	obj_type := reflect.TypeOf(obj).Elem()
	field_names := make([]string, 0)
	value_mappings := make([]string, 0)
	for i := 0; i < obj_type.NumField(); i++ {
		field := obj_type.Field(i)
		tag := field.Tag.Get("db")
		if tag != obj.IdFieldName() {
			field_names = append(field_names, tag)
			value_mappings = append(value_mappings, ":"+tag)
		}
	}
	cmd := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(field_names, ","),
		strings.Join(value_mappings, ","))

	return db.Tx(func(tx *sqlx.Tx) error {
		r, err := tx.NamedExec(cmd, obj)
		if err != nil {
			return errors.Wrapf(err, "Error while exec SQL\n%s\n", cmd)
		}
		id, _ := r.LastInsertId()
		obj.SetId(id)
		return nil
	})
}

func (db *Database) Get(obj Mapper, id int64) error {
	table := obj.TableName()
	pk_name := obj.IdFieldName()
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", table, pk_name, id)
	return errors.Wrapf(db.db.Get(obj, cmd), "Error whlie select\n%s\n", cmd)
}
