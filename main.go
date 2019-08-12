package main

import (
	"os"
	"path/filepath"

	"github.com/naoty/ttodo/todo"
	"github.com/naoty/ttodo/views"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// To reset background color of previous selected rows
	// https://github.com/rivo/tview/issues/270
	app.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	pages := tview.NewPages()
	pages.Box.SetBackgroundColor(tcell.ColorDefault)

	dir := os.Getenv("TODO_PATH")
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	todoPath := filepath.Join(dir, ".todo.json")
	todos, err := todo.LoadTodos(todoPath)

	if err != nil {
		panic(err)
	}

	table := views.NewTable()

	for _, todo := range todos {
		table.AddTodoRow(todo)
	}

	descriptionView := views.NewDescription()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)
	flex.Box.SetBackgroundColor(tcell.ColorDefault)

	table.SetSelectedFunc(func(row, column int) {
		if descriptionView.IsShown {
			descriptionView.IsShown = false
			flex.RemoveItem(descriptionView)
		} else {
			descriptionView.IsShown = true

			if row >= 1 && row <= len(todos) {
				todo := todos[row-1]
				descriptionView.SetText(todo.Description)
				flex.AddItem(descriptionView, 0, 1, true)
			}
		}
	})

	table.SetSelectionChangedFunc(func(row, column int) {
		if !descriptionView.IsShown {
			return
		}

		if row >= 1 && row <= len(todos) {
			todo := todos[row-1]
			descriptionView.SetText(todo.Description)
		}
	})

	form := views.NewForm(func(td todo.Todo) {
		err := todo.SaveTodo(td, todoPath)

		if err != nil {
			panic(err)
		}

		todos, err = todo.LoadTodos(todoPath)

		if err != nil {
			panic(err)
		}

		table.UpdateTodoRows(todos)
		pages.SwitchToPage("main")
	}, func() {
		pages.SwitchToPage("main")
	})

	pages.
		AddPage("main", flex, true, true).
		AddPage("new", form, true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			pages.SwitchToPage("new")
			return nil
		case 'q':
			app.Stop()
			return nil
		default:
			return event
		}
	})

	err = app.SetRoot(pages, true).Run()

	if err != nil {
		panic(err)
	}
}
