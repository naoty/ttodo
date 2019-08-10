package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetSelectable(true, false).
		SetSelectedStyle(tcell.ColorDefault, tcell.Color100, 0)

	table.SetCell(0, 0, tview.NewTableCell("Done").SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Deadline").SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Assignee").SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("Title").SetSelectable(false))
	table.SetCellSimple(1, 0, "[ ]")
	table.SetCellSimple(1, 1, "2019-08-10")
	table.SetCellSimple(1, 2, "naoty")
	table.SetCellSimple(1, 3, "TODOのリストを表示する")
	table.SetCellSimple(2, 0, "[ ]")
	table.SetCellSimple(2, 1, "2019-08-10")
	table.SetCellSimple(2, 2, "naoty")
	table.SetCellSimple(2, 3, "TODOの詳細を表示する")

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
