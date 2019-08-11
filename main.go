package main

import (
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

	table := tview.NewTable().
		SetSelectable(true, false).
		SetSelectedStyle(tcell.ColorDefault, tcell.Color100, 0)
	table.Box.SetBackgroundColor(tcell.ColorDefault)

	table.SetCell(0, 0, tview.NewTableCell("Done").SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Deadline").SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Assignee").SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("Title").SetSelectable(false).SetExpansion(1))
	table.SetCellSimple(1, 0, "[ ]")
	table.SetCellSimple(1, 1, "2019-08-10")
	table.SetCellSimple(1, 2, "naoty")
	table.SetCellSimple(1, 3, "TODOのリストを表示する")
	table.SetCellSimple(2, 0, "[ ]")
	table.SetCellSimple(2, 1, "2019-08-10")
	table.SetCellSimple(2, 2, "naoty")
	table.SetCellSimple(2, 3, "TODOの詳細を表示する")

	border := tview.NewTextView().SetText("Description")
	border.Box.SetBackgroundColor(tcell.Color32)

	textView := tview.NewTextView().SetText("これはダミーです。")
	textView.Box.SetBackgroundColor(tcell.ColorDefault)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true).
		AddItem(border, 1, 0, false).
		AddItem(textView, 0, 1, false)
	flex.Box.SetBackgroundColor(tcell.ColorDefault)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
