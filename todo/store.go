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
	todos       map[uuid.UUID]Todo
	firstTodoID *uuid.UUID
	lastTodoID  *uuid.UUID
	source      io.ReadWriteSeeker
	subscribers []chan []Todo
}

// NewStore initializes a singleton Store once.
func NewStore(source io.ReadWriteSeeker) *Store {
	once.Do(func() {
		singletonStore = &Store{
			todos:       map[uuid.UUID]Todo{},
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

	var todos []Todo
	err = json.Unmarshal(data, &todos)

	if err != nil {
		return err
	}

	_, err = store.source.Seek(0, io.SeekStart)

	if err != nil {
		return err
	}

	for _, todo := range todos {
		store.todos[todo.ID] = todo

		if todo.PreviousID == nil {
			id := todo.ID
			store.firstTodoID = &id
		}
	}

	store.publish()

	return nil
}

// SaveTodos saves todos into source.
func (store *Store) SaveTodos() error {
	todos := store.todosList()

	indent := strings.Repeat(" ", 4)
	data, err := json.MarshalIndent(todos, "", indent)

	if err != nil {
		return err
	}

	_, err = store.source.Write(data)

	if err != nil {
		return err
	}

	_, err = store.source.Seek(0, io.SeekStart)

	if err != nil {
		return err
	}

	return nil
}

// AppendTodo saves a Todo with given parameters.
func (store *Store) AppendTodo(title, description, assignee string, deadline *time.Time) {
	todo := Todo{
		ID:          uuid.New(),
		PreviousID:  store.lastTodoID,
		Title:       title,
		Description: description,
		Assignee:    assignee,
		Deadline:    deadline,
		Done:        false,
	}

	if store.lastTodoID != nil {
		lastTodo := store.todos[*store.lastTodoID]
		lastTodo.NextID = &todo.ID
		store.todos[*store.lastTodoID] = lastTodo
	}

	if store.firstTodoID == nil {
		store.firstTodoID = &todo.ID
	}

	store.lastTodoID = &todo.ID

	store.todos[todo.ID] = todo
	store.publish()
	store.SaveTodos()
}

// ToggleDone toggles done of the todo with given id.
func (store *Store) ToggleDone(id uuid.UUID) {
	todo, ok := store.todos[id]

	if !ok {
		return
	}

	todo.Done = !todo.Done
	store.todos[id] = todo
	store.publish()
	store.SaveTodos()
}

func (store *Store) publish() {
	todos := store.todosList()

	go func() {
		for _, subscriber := range store.subscribers {
			subscriber <- todos
		}
	}()
}

func (store *Store) todosList() []Todo {
	todos := make([]Todo, len(store.todos))

	i := 0
	nextID := store.firstTodoID

	for {
		if nextID == nil {
			break
		}

		todo := store.todos[*nextID]
		todos[i] = todo
		i++
		nextID = todo.NextID
	}

	return todos
}
