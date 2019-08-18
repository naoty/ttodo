package todo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var singletonStore *Store
var once sync.Once

// Store is the single source of truth for todos.
type Store struct {
	todos       []Todo
	source      io.ReadWriteSeeker
	subscribers []chan []Todo
}

// NewStore initializes a singleton Store once.
func NewStore(source io.ReadWriteSeeker) *Store {
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

// UnregisterAll closes all subscribers.
func (store *Store) UnregisterAll() {
	for _, subscriber := range store.subscribers {
		close(subscriber)
	}
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

	_, err = store.source.Seek(0, io.SeekStart)

	if err != nil {
		return err
	}

	store.publish()

	return nil
}

// SaveTodos saves todos into source.
func (store *Store) SaveTodos() error {
	indent := strings.Repeat(" ", 4)
	data, err := json.MarshalIndent(store.todos, "", indent)

	if err != nil {
		return err
	}

	_, err = store.source.Write(data)

	if err != nil {
		return err
	}

	return nil
}

// AppendTodo saves a Todo with given parameters.
func (store *Store) AppendTodo(title, description, assignee string, deadline *time.Time) {
	todo := Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Assignee:    assignee,
		Deadline:    deadline,
		Done:        false,
	}
	store.todos = append(store.todos, todo)
	store.publish()
	store.SaveTodos()
}

func (store *Store) publish() {
	go func() {
		for _, subscriber := range store.subscribers {
			subscriber <- store.todos
		}
	}()
}
