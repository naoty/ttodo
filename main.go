package main

import (
	"time"

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

	jst, _ := time.LoadLocation("Asia/Tokyo")
	todos := []todo{
		todo{
			title:       "TODOのリストを表示する",
			description: "TODOのリストを表形式で表示したい",
			assignee:    "naoty",
			deadline:    time.Date(2019, 8, 10, 0, 0, 0, 0, jst),
			done:        true,
		},
		todo{
			title:       "TODOの詳細を表示する",
			description: "TODOの詳細を表の下に表示したい。できればEnterで表示/非表示を切り替えたい",
			assignee:    "naoty",
			deadline:    time.Date(2019, 8, 11, 0, 0, 0, 0, jst),
			done:        false,
		},
	}

	for i, todo := range todos {
		if todo.done {
			table.SetCellSimple(i+1, 0, tview.Escape("[x]"))
		} else {
			table.SetCellSimple(i+1, 0, "[ ]")
		}
		table.SetCellSimple(i+1, 1, todo.deadline.Format("2006-01-02"))
		table.SetCellSimple(i+1, 2, todo.assignee)
		table.SetCellSimple(i+1, 3, todo.title)
	}

	border := tview.NewTextView().SetText("Description")
	border.Box.SetBackgroundColor(tcell.Color32)

	textView := tview.NewTextView()
	textView.Box.SetBackgroundColor(tcell.ColorDefault)

	descriptionView := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(border, 1, 0, false).
		AddItem(textView, 0, 1, false)
	descriptionView.Box.SetBackgroundColor(tcell.ColorDefault)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)
	flex.Box.SetBackgroundColor(tcell.ColorDefault)

	descriptionViewShown := false
	table.SetSelectedFunc(func(row, column int) {
		if descriptionViewShown {
			descriptionViewShown = false
			flex.RemoveItem(descriptionView)
		} else {
			descriptionViewShown = true

			if row >= 1 && row <= len(todos) {
				todo := todos[row-1]
				textView.SetText(todo.description)
				flex.AddItem(descriptionView, 0, 1, true)
			}
		}
	})

	table.SetSelectionChangedFunc(func(row, column int) {
		if !descriptionViewShown {
			return
		}

		if row >= 1 && row <= len(todos) {
			todo := todos[row-1]
			textView.SetText(todo.description)
		}
	})

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
