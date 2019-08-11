package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Todo represents a TODO.
type Todo struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Assignee    string     `json:"assignee"`
	Deadline    *time.Time `json:"deadline"`
	Done        bool       `json:"done"`
}

// LoadTodos reads file at given path and decodes the content into []Todo.
func LoadTodos(path string) ([]Todo, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to read data: %v", err)
	}

	var todos []Todo
	err = json.Unmarshal(contents, &todos)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode data: %v", err)
	}

	return todos, nil
}
