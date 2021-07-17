package tree

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Todo struct {
	id         int64
	title      string
	status     string
	hidden     bool
	created_at int64
	updated_at int64
}

func (todo *Todo) SetId(id int64) {
	todo.id = id
}

func (todo *Todo) GetTableName() string {
	return "todo"
}

func (todo *Todo) GetFieldNames() string {
	return "title, status, hidden, created_at, updated_at"
}

func (todo *Todo) GetPkFieldName() string {
	return "todo_id"
}

func (todo *Todo) GetValueList() string {
	hiddenInt := 0
	if todo.hidden {
		hiddenInt = 1
	}
	return fmt.Sprintf("'%s', '%s', %d, %d, %d",
		todo.title,
		todo.status,
		hiddenInt,
		todo.created_at,
		todo.updated_at)
}

func (todo *Todo) GetUpdateList(fields []string) string {
	// TEMP
	return ""
}

func (todo *Todo) Deserialize(row *sql.Row) error {
	err := row.Scan(
		&todo.id,
		&todo.title,
		&todo.status,
		&todo.hidden,
		&todo.created_at,
		&todo.updated_at)
	if err != nil {
		return errors.Wrapf(err, "Error while scanning row:")
	}
	return nil
}

func CreateTodo(title string) *Todo {
	now := time.Now().Unix()
	return &Todo{
		id:         -1,
		title:      title,
		status:     "Not Started",
		hidden:     false,
		created_at: now,
		updated_at: now,
	}
}
