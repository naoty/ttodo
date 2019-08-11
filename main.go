package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

	table := tview.NewTable().
		SetSelectable(true, false).
		SetSelectedStyle(tcell.ColorDefault, tcell.Color100, 0)
	table.Box.SetBackgroundColor(tcell.ColorDefault)

	table.SetCell(0, 0, tview.NewTableCell("Done").SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Deadline").SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Assignee").SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("Title").SetSelectable(false).SetExpansion(1))

	todos, err := loadTodos()

	if err != nil {
		panic(err)
	}

	for i, todo := range todos {
		if todo.Done {
			table.SetCellSimple(i+1, 0, tview.Escape("[x]"))
		} else {
			table.SetCellSimple(i+1, 0, "[ ]")
		}
		if todo.Deadline == nil {
			table.SetCellSimple(i+1, 1, "")
		} else {
			table.SetCellSimple(i+1, 1, (*todo.Deadline).Format("2006-01-02"))
		}
		table.SetCellSimple(i+1, 2, todo.Assignee)
		table.SetCellSimple(i+1, 3, todo.Title)
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

	form := tview.NewForm().
		AddInputField("Title", "", 0, nil, nil).
		AddInputField("Deadline", "", 0, nil, nil).
		AddInputField("Assignee", "", 0, nil, nil).
		AddInputField("Description", "", 0, nil, nil).
		AddButton("Save", func() {
			// TODO: save a TODO
			pages.SwitchToPage("main")
		}).
		AddButton("Quit", func() {
			pages.SwitchToPage("main")
		})
	form.SetBackgroundColor(tcell.ColorDefault)

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

func loadTodos() ([]Todo, error) {
	dir := os.Getenv("TODO_PATH")
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	path := filepath.Join(dir, ".todo.json")
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
