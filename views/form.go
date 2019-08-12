package views

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Form represents a form to save a TODO.
type Form struct {
	*tview.Form
}

// NewForm returns a new Form.
func NewForm(saveHandler func(td todo.Todo), quitHandler func()) *Form {
	form := tview.NewForm()
	form.SetBackgroundColor(tcell.ColorDefault)

	titleInput := tview.NewInputField().SetLabel("Title")
	form.AddFormItem(titleInput)

	deadlineInput := tview.NewInputField().SetLabel("Deadline")
	form.AddFormItem(deadlineInput)

	assigneeInput := tview.NewInputField().SetLabel("Assignee")
	form.AddFormItem(assigneeInput)

	descriptionInput := tview.NewInputField().SetLabel("Description")
	form.AddFormItem(descriptionInput)

	form.AddButton("Save", func() {
		deadline, _ := time.Parse("2006-01-02", deadlineInput.GetText())

		td := todo.Todo{
			Title:       titleInput.GetText(),
			Description: descriptionInput.GetText(),
			Done:        false,
			Deadline:    &deadline,
			Assignee:    assigneeInput.GetText(),
		}
		saveHandler(td)
	})
	form.AddButton("Quit", quitHandler)

	return &Form{form}
}
