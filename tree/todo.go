package tree

import (
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

func (todo *Todo) TableName() string {
	return "todo"
}

func (todo *Todo) IdFieldName() string {
	return "todo_id"
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
