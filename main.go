package main

import "github.com/rivo/tview"

func main() {
	LoadApp()
}

//LoadApp will load the app
func LoadApp() {
	app := tview.NewApplication()

	//heightField := tview.NewInputField().SetLabel("Height")

	form := tview.NewForm().
		AddInputField("height", "", 10, nil, nil).
		AddInputField("Weight", "", 10, nil, nil).
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

//Add an event listener to the calculate button, if i select the calculate button it should open a new page
