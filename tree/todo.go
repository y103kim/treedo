package tree

import (
	"time"
)

type Todo struct {
	Id        int64  `db:"todo_id"`
	Title     string `db:"title"`
	Status    string `db:"status"`
	Hidden    int64  `db:"hidden"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (todo *Todo) GetId() int64 {
	return todo.Id
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
		Hidden:    0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
