package tree

import (
	"fmt"
	"time"
)

type Node struct {
	id         int64
	title      string
	status     string
	hidden     bool
	created_at time.Time
	updated_at time.Time
}

func (node *Node) SetId(id int64) {
	node.id = id
}

func (node *Node) GetTableName() string {
	return "node"
}

func (node *Node) GetFieldNames() string {
	return "title, status, hidden"
}

func (node *Node) GetValueList() string {
	hiddenInt := 0
	if node.hidden {
		hiddenInt = 1
	}
	return fmt.Sprintf("%s, %s, %d", node.title, node.status, hiddenInt)
}

func (node *Node) GetUpdateList(fields []string) string {
	// TEMP
	return ""
}

func (node *Node) Deserialize(db_output string) error {
	// TEMP
	return nil
}

func Create(title string) *Node {
	return &Node{
		id:         -1,
		title:      title,
		status:     "Not Started",
		hidden:     false,
		created_at: time.Now(),
		updated_at: time.Now(),
	}
}
