package tree

import (
	"fmt"
	"time"
)

type Todo struct {
	id         int64
	title      string
	status     string
	hidden     bool
	created_at time.Time
	updated_at time.Time
}

func (todo *Todo) SetId(id int64) {
	todo.id = id
}

func (todo *Todo) GetTableName() string {
	return "todo"
}

func (todo *Todo) GetFieldNames() string {
	return "title, status, hidden"
}

func (todo *Todo) GetValueList() string {
	hiddenInt := 0
	if todo.hidden {
		hiddenInt = 1
	}
	return fmt.Sprintf("'%s', '%s', '%d'", todo.title, todo.status, hiddenInt)
}

func (todo *Todo) GetUpdateList(fields []string) string {
	// TEMP
	return ""
}

func (todo *Todo) Deserialize(db_output string) error {
	// TEMP
	return nil
}

func Create(title string) *Todo {
	return &Todo{
		id:         -1,
		title:      title,
		status:     "Not Started",
		hidden:     false,
		created_at: time.Now(),
		updated_at: time.Now(),
	}
}
