package todo

import (
	"time"

	"github.com/google/uuid"
)

// Todo represents a TODO.
type Todo struct {
	ID          uuid.UUID  `json:"id"`
	NextID      *uuid.UUID `json:"nextId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Assignee    string     `json:"assignee"`
	Deadline    *time.Time `json:"deadline"`
	Done        bool       `json:"done"`
}
