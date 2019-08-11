package views

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Description represents a view to display the description of a TODO.
type Description struct {
	*tview.Flex
	IsShown bool

	body *tview.TextView
}

// NewDescription returns a new Description.
func NewDescription() *Description {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBackgroundColor(tcell.ColorDefault)

	header := tview.NewTextView().SetText("Description")
	header.SetBackgroundColor(tcell.Color32)
	flex.AddItem(header, 1, 0, false)

	body := tview.NewTextView()
	body.SetBackgroundColor(tcell.ColorDefault)
	flex.AddItem(body, 0, 1, false)

	return &Description{Flex: flex, IsShown: false, body: body}
}

// SetText sets given text to body.
func (d *Description) SetText(text string) {
	d.body.SetText(text)
}
