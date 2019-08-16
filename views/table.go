package views

import (
	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Table represents a table to list TODOs.
type Table struct {
	*tview.Table

	todos []todo.Todo
}

// NewTable returns a new Table.
func NewTable() *Table {
	table := tview.NewTable().
		SetSelectable(true, false).
		SetCell(0, 0, tview.NewTableCell("Done").SetSelectable(false)).
		SetCell(0, 1, tview.NewTableCell("Deadline").SetSelectable(false)).
		SetCell(0, 2, tview.NewTableCell("Assignee").SetSelectable(false)).
		SetCell(0, 3, tview.NewTableCell("Title").SetSelectable(false).SetExpansion(1))

	return &Table{
		Table: table,
		todos: []todo.Todo{},
	}
}

// SetSelectedFunc sets wrapper function to pass selected todo to original
// function.
func (table *Table) SetSelectedFunc(f func(todo.Todo)) {
	table.Table.SetSelectedFunc(func(row, column int) {
		if row > 0 && row <= len(table.todos) {
			todo := table.todos[row-1]
			f(todo)
		}
	})
}

// SetSelectionChangedFunc sets wrapper function to pass selected todo to
// original function.
func (table *Table) SetSelectionChangedFunc(f func(todo todo.Todo)) {
	table.Table.SetSelectionChangedFunc(func(row, column int) {
		if row > 0 && row <= len(table.todos) {
			todo := table.todos[row-1]
			f(todo)
		}
	})
}

// Update updates rows along with given todos.
func (table *Table) Update(todos []todo.Todo) {
	table.todos = todos

	for row := table.GetRowCount(); row > 0; row-- {
		table.RemoveRow(row)
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
}
