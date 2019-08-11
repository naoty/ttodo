package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
			return nil
		default:
			return event
		}
	})

	// To reset background color of previous selected rows
	// https://github.com/rivo/tview/issues/270
	app.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

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

	border := tview.NewTextView().SetText("Description")
	border.Box.SetBackgroundColor(tcell.Color32)

	textView := tview.NewTextView()
	textView.Box.SetBackgroundColor(tcell.ColorDefault)

	DescriptionView := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(border, 1, 0, false).
		AddItem(textView, 0, 1, false)
	DescriptionView.Box.SetBackgroundColor(tcell.ColorDefault)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)
	flex.Box.SetBackgroundColor(tcell.ColorDefault)

	DescriptionViewShown := false
	table.SetSelectedFunc(func(row, column int) {
		if DescriptionViewShown {
			DescriptionViewShown = false
			flex.RemoveItem(DescriptionView)
		} else {
			DescriptionViewShown = true

			if row >= 1 && row <= len(todos) {
				todo := todos[row-1]
				textView.SetText(todo.Description)
				flex.AddItem(DescriptionView, 0, 1, true)
			}
		}
	})

	table.SetSelectionChangedFunc(func(row, column int) {
		if !DescriptionViewShown {
			return
		}

		if row >= 1 && row <= len(todos) {
			todo := todos[row-1]
			textView.SetText(todo.Description)
		}
	})

	err = app.SetRoot(flex, true).SetFocus(flex).Run()

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
