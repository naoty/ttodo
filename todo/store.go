package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
)

var singletonStore *Store
var once sync.Once

// Store is the single source of truth for todos.
type Store struct {
	todos  []Todo
	source io.Reader
}

// GetStore initializes a singleton Store once.
func GetStore(source io.Reader) *Store {
	once.Do(func() {
		singletonStore = &Store{
			todos:  []Todo{},
			source: source,
		}
	})
	return singletonStore
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

	// DEBUG
	fmt.Println(store.todos)

	return nil
}
