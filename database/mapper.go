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
	GetId() int64
	TableName() string
	IdFieldName() string
}

func (db *Database) Insert(mapper Mapper) error {
	table := mapper.TableName()
	mapper_type := reflect.TypeOf(mapper).Elem()
	field_names := make([]string, 0)
	value_mappings := make([]string, 0)
	for i := 0; i < mapper_type.NumField(); i++ {
		field := mapper_type.Field(i)
		tag := field.Tag.Get("db")
		if tag != mapper.IdFieldName() {
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
		r, err := tx.NamedExec(cmd, mapper)
		if err != nil {
			return errors.Wrapf(err, "Error while exec SQL\n%s\n", cmd)
		}
		id, _ := r.LastInsertId()
		mapper.SetId(id)
		return nil
	})
}

func (db *Database) Get(mapper Mapper, id int64) error {
	table := mapper.TableName()
	pk_name := mapper.IdFieldName()
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", table, pk_name, id)
	return errors.Wrapf(db.db.Get(mapper, cmd), "Error whlie select\n%s\n", cmd)
}
