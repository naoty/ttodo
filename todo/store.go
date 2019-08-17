package todo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
	"time"
)

var singletonStore *Store
var once sync.Once

// Store is the single source of truth for todos.
type Store struct {
	todos       []Todo
	source      io.Reader
	subscribers []chan []Todo
}

// NewStore initializes a singleton Store once.
func NewStore(source io.Reader) *Store {
	once.Do(func() {
		singletonStore = &Store{
			todos:       []Todo{},
			source:      source,
			subscribers: []chan []Todo{},
		}
	})
	return singletonStore
}

// GetStore returns a singleton store.
func GetStore() *Store {
	return singletonStore
}

// Register sets given subscriber to store.subscriber.
func (store *Store) Register(subscriber chan []Todo) {
	store.subscribers = append(store.subscribers, subscriber)
}

// LoadTodos loads todos from source.
func (store *Store) LoadTodos() error {
	data, err := ioutil.ReadAll(store.source)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &store.todos)

	if err != nil {
		return err
	}

	store.publish()

	return nil
}

// AppendTodo saves a Todo with given parameters.
func (store *Store) AppendTodo(title, description, assignee string, deadline *time.Time) {
	todo := Todo{
		Title:       title,
		Description: description,
		Assignee:    assignee,
		Deadline:    deadline,
		Done:        false,
	}
	store.todos = append(store.todos, todo)
	store.publish()
}

func (store *Store) publish() {
	go func() {
		for _, subscriber := range store.subscribers {
			subscriber <- store.todos
		}
	}()
}
