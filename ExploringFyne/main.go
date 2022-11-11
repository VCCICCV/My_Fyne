package main

// PROJECT_NAME:Fyne
// DATE:2022/10/29 14:26
// USER:21005
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Hello World")
	window.Resize(fyne.NewSize(600, 500))
	label := widget.NewLabel("New Content")
	vBox := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		label,
		layout.NewSpacer())
	window.SetContent(vBox)
	window.ShowAndRun()
}
