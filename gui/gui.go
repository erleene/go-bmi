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
		AddButton(calculateButton.GetLabel(), func() {
			calc.Calculate(heightField.GetText(), weightField.GetText())
		}).
		AddButton(exitButton.GetLabel(), func() {
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

//it has to be a type func()
//so i need a method to return a func()
//this func should take in parameters and return a func
