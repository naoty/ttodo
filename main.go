package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	// Keep background color as users set
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault

	app := tview.NewApplication()
	table := tview.NewTable()

	table.SetCellSimple(0, 0, "[ ]")
	table.SetCellSimple(0, 1, "2019-08-10")
	table.SetCellSimple(0, 2, "naoty")
	table.SetCellSimple(0, 3, "go moduleの使い方を思い出す")
	table.SetCellSimple(1, 0, "[ ]")
	table.SetCellSimple(1, 1, "2019-08-10")
	table.SetCellSimple(1, 2, "naoty")
	table.SetCellSimple(1, 3, "VSCodeのGoの設定を見直す")
	table.SetCellSimple(2, 0, "[ ]")
	table.SetCellSimple(2, 1, "2019-08-11")
	table.SetCellSimple(2, 2, "naoty")
	table.SetCellSimple(2, 3, "rivo/tviewの使い方を理解する")

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
