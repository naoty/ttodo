package todo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
)

var singletonStore *Store
var once sync.Once

// Store is the single source of truth for todos.
type Store struct {
	todos       []Todo
	source      io.Reader
	subscribers []chan []Todo
}

// GetStore initializes a singleton Store once.
func GetStore(source io.Reader) *Store {
	once.Do(func() {
		singletonStore = &Store{
			todos:       []Todo{},
			source:      source,
			subscribers: []chan []Todo{},
		}
	})
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

func (store *Store) publish() {
	go func() {
		for _, subscriber := range store.subscribers {
			subscriber <- store.todos
		}
	}()
}
