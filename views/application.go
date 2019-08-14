package views

import (
	"github.com/rivo/tview"
)

// Application represents whole TUI application.
type Application struct {
	*tview.Application
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

	return &Application{app}
}
