package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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

// SaveTodos encodes todos and writes them into file at given path.
func SaveTodos(todos []Todo, path string) error {
	indent := strings.Repeat(" ", 4)
	json, err := json.MarshalIndent(todos, "", indent)

	if err != nil {
		return fmt.Errorf("Failed to encode data: %v", err)
	}

	err = ioutil.WriteFile(path, json, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write data: %v", err)
	}

	return nil
}

// SaveTodo saves a TODO into file at given path.
func SaveTodo(td Todo, path string) error {
	todos, err := LoadTodos(path)

	if err != nil {
		return err
	}

	todos = append(todos, td)
	err = SaveTodos(todos, path)

	if err != nil {
		return err
	}

	return nil
}
