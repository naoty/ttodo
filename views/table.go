package views

import (
	"github.com/gdamore/tcell"
	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Table represents a table to list TODOs.
type Table struct {
	*tview.Table
}

// NewTable returns a new Table.
func NewTable() *Table {
	table := tview.NewTable().
		SetSelectable(true, false).
		SetSelectedStyle(tcell.ColorDefault, tcell.Color100, 0)

	done := tview.NewTableCell("Done").SetSelectable(false)
	table.SetCell(0, 0, done)

	deadline := tview.NewTableCell("Deadline").SetSelectable(false)
	table.SetCell(0, 1, deadline)

	assignee := tview.NewTableCell("Assignee").SetSelectable(false)
	table.SetCell(0, 2, assignee)

	title := tview.NewTableCell("Title").SetSelectable(false).SetExpansion(1)
	table.SetCell(0, 3, title)

	table.SetBackgroundColor(tcell.ColorDefault)

	return &Table{table}
}

// AddTodoRow adds a row into table for a given TODO.
func (t *Table) AddTodoRow(td todo.Todo) {
	row := t.GetRowCount()

	if td.Done {
		t.SetCellSimple(row, 0, tview.Escape("[x]"))
	} else {
		t.SetCellSimple(row, 0, "[ ]")
	}
	if td.Deadline == nil {
		t.SetCellSimple(row, 1, "")
	} else {
		t.SetCellSimple(row, 1, (*td.Deadline).Format("2006-01-02"))
	}
	t.SetCellSimple(row, 2, td.Assignee)
	t.SetCellSimple(row, 3, td.Title)
}

// UpdateTodoRows update rows to display todos.
func (t *Table) UpdateTodoRows(todos []todo.Todo) {
	total := t.GetRowCount()

	// Remove rows except for headers
	for row := total - 1; row > 0; row-- {
		t.RemoveRow(row)
	}

	for _, todo := range todos {
		t.AddTodoRow(todo)
	}
}
