package gui

import (
	"github.com/erleene/go-bmi/calc"
	"github.com/rivo/tview"
)

//LoadApp will load the app
func LoadApp() {

	app := tview.NewApplication()

	//input fields with labels
	heightField := tview.NewInputField().SetLabel("Height")
	weightField := tview.NewInputField().SetLabel("Weight")

	//buttons for calculating and exiting app
	calculateButton := tview.NewButton("Calculate BMI")
	exitButton := tview.NewButton("Exit")

	form := tview.NewForm().
		AddInputField(heightField.GetLabel(), "", 10, nil, nil).
		AddInputField(weightField.GetLabel(), "", 10, nil, nil).
		AddButton("Calculate BMI", calc.calculate(heightField.GetText(), weightField.GetText())).
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
