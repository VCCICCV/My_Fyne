package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PROJECT_NAME:Fyne
// DATE:2022/10/29 21:29
// USER:21005

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("我的APP")

	output, entry, btn := myApp.makeUI()
	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}
func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("看看你的")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	// 重要性
	btn.Importance = widget.HighImportance
	app.output = output
	return output, entry, btn
}
