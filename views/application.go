package views

import (
	"github.com/gdamore/tcell"
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
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor: tcell.ColorDefault,
		ContrastBackgroundColor:  tcell.ColorBlue,
		PrimaryTextColor:         tcell.ColorDefault,
		SecondaryTextColor:       tcell.ColorDefault,
	}

	table := NewTable()

	description := NewDescription()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)

	table.SetSelectedFunc(func(todo todo.Todo) {
		if description.IsShown {
			description.IsShown = false
			flex.RemoveItem(description)
		} else {
			description.IsShown = true
			description.SetText(todo.Description)
			flex.AddItem(description, 0, 1, false)
		}
	})

	table.SetSelectionChangedFunc(func(todo todo.Todo) {
		if description.IsShown {
			description.SetText(todo.Description)
		}
	})

	pages := tview.NewPages().
		AddPage("list", flex, true, true)

	form := NewForm(func() {
		pages.SwitchToPage("list")
	})
	pages.AddPage("form", form, true, false)

	app := tview.NewApplication().
		SetRoot(pages, true)

	// https://github.com/rivo/tview/issues/270
	app.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			pages.SwitchToPage("form")
			return nil
		default:
			return event
		}
	})

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
