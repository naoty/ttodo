package views

import (
	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Application represents whole TUI application.
type Application struct {
	*tview.Application

	table      *Table
	subscriber chan []todo.Todo
}

// NewApplication initializes and returns Application.
func NewApplication() *Application {
	table := NewTable()

	app := tview.NewApplication().
		SetRoot(table, true)

	return &Application{
		Application: app,
		table:       table,
		subscriber:  make(chan []todo.Todo),
	}
}

// Run starts TUI application.
func (app *Application) Run() error {
	go func() {
		for {
			todos := <-app.subscriber
			app.Update(todos)
		}
	}()

	return app.Application.Run()
}

// Subscribe subscribes to the update of todos.
func (app *Application) Subscribe(store *todo.Store) {
	store.Register(app.subscriber)
}

// Update updates views along with given todos.
func (app *Application) Update(todos []todo.Todo) {
	app.table.Update(todos)
}
