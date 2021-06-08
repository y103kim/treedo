package tree

import "time"

type Node struct {
	id         int
	title      string
	status     string
	hidden     bool
	created_at time.Time
	updated_at time.Time
}
