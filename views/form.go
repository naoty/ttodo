package views

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Form represents a form to save a TODO.
type Form struct {
	*tview.Form
}

// NewForm returns a new Form.
func NewForm(saveHandler, quitHandler func()) *Form {
	form := tview.NewForm().
		AddInputField("Title", "", 0, nil, nil).
		AddInputField("Deadline", "", 0, nil, nil).
		AddInputField("Assignee", "", 0, nil, nil).
		AddInputField("Description", "", 0, nil, nil).
		AddButton("Save", saveHandler).
		AddButton("Quit", quitHandler)

	form.SetBackgroundColor(tcell.ColorDefault)

	return &Form{form}
}
