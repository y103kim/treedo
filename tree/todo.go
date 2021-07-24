package tree

import (
	"fmt"
	"time"
)

type Todo struct {
	Id        int64  `db:"todo_id"`
	Title     string `db:"title"`
	Status    string `db:"status"`
	Hidden    bool   `db:"hidden"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (todo *Todo) SetId(id int64) {
	todo.Id = id
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
	if todo.Hidden {
		hiddenInt = 1
	}
	return fmt.Sprintf("'%s', '%s', %d, %d, %d",
		todo.Title,
		todo.Status,
		hiddenInt,
		todo.CreatedAt,
		todo.UpdatedAt)
}

func (todo *Todo) GetUpdateList(fields []string) string {
	// TEMP
	return ""
}

func CreateTodo(title string) *Todo {
	now := time.Now().Unix()
	return &Todo{
		Id:        -1,
		Title:     title,
		Status:    "Not Started",
		Hidden:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
