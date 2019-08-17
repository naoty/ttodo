package views

import (
	"time"

	"github.com/naoty/ttodo/todo"
	"github.com/rivo/tview"
)

// Form represents a form to save a TODO.
type Form struct {
	*tview.Form
}

// NewForm returns a new Form.
func NewForm(handler func()) *Form {
	form := tview.NewForm()
	titleInput := tview.NewInputField().SetLabel("Title")
	deadlineInput := tview.NewInputField().SetLabel("Deadline")
	assigneeInput := tview.NewInputField().SetLabel("Assignee")
	descriptionInput := tview.NewInputField().SetLabel("Description")

	form.
		AddFormItem(titleInput).
		AddFormItem(deadlineInput).
		AddFormItem(assigneeInput).
		AddFormItem(descriptionInput).
		AddButton("Save", func() {
			title := titleInput.GetText()

			if title == "" {
				return
			}

			deadline, _ := time.Parse("2006-01-02", deadlineInput.GetText())

			todo.GetStore().AppendTodo(
				title,
				descriptionInput.GetText(),
				assigneeInput.GetText(),
				&deadline,
			)

			handler()
		}).
		AddButton("Quit", func() {
			handler()
		})

	return &Form{form}
}
