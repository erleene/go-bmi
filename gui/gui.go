package gui

import "github.com/rivo/tview"

//LoadApp will load the app
func LoadApp() {

	app := tview.NewApplication()

	heightField := tview.NewInputField().SetLabel("Height")
	weightField := tview.NewInputField().SetLabel("Weight")

	form := tview.NewForm().
		AddInputField(heightField.GetText(), "", 10, nil, nil).
		AddInputField(weightField.GetText(), "", 10, nil, nil).
		AddButton("Calculate BMI", nil).
		AddButton("Exit", func() {
			app.Stop()
		}).
		SetCancelFunc(func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("BMI Calculator").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
