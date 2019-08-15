package views

import (
	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Application represents whole TUI application.
type Application struct {
	*tview.Application

	table      *tview.Table
	subscriber chan []todo.Todo
}

// NewApplication initializes and returns Application.
func NewApplication() *Application {
	table := tview.NewTable().
		SetSelectable(true, false).
		SetCell(0, 0, tview.NewTableCell("Done").SetSelectable(false)).
		SetCell(0, 1, tview.NewTableCell("Deadline").SetSelectable(false)).
		SetCell(0, 2, tview.NewTableCell("Assignee").SetSelectable(false)).
		SetCell(0, 3, tview.NewTableCell("Title").SetSelectable(false).SetExpansion(1))

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
	for row := app.table.GetRowCount(); row > 0; row-- {
		app.table.RemoveRow(row)
	}

	for i, todo := range todos {
		if todo.Done {
			app.table.SetCellSimple(i+1, 0, tview.Escape("[x]"))
		} else {
			app.table.SetCellSimple(i+1, 0, tview.Escape("[ ]"))
		}
		if todo.Deadline == nil {
			app.table.SetCellSimple(i+1, 1, "")
		} else {
			app.table.SetCellSimple(i+1, 1, (*todo.Deadline).Format("2006-01-02"))
		}
		app.table.SetCellSimple(i+1, 2, todo.Assignee)
		app.table.SetCellSimple(i+1, 3, todo.Title)
	}
}
