package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"os"
)

// PROJECT_NAME:Fyne
// DATE:2022/11/10 9:45
// USER:21005
// import (
//
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/widget"
//	"github.com/flopp/go-findfont"
//	"github.com/golang/freetype/truetype"
//	"os"
//
// )
func init() {
	fontPath, err := findfont.Find("SIMYOU.TTF")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found 'arial.ttf' in '%s'\n", fontPath)
	// load the font with the freetype library
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

//func main() {
//	myApp := app.New()
//	myWindow := myApp.NewWindow("Widget")
//	myWindow.Resize(fyne.NewSize(600, 500))
//	label := widget.NewLabel("New Content")
//	vBox := container.New(
//		layout.NewVBoxLayout(),
//		layout.NewSpacer(),
//		label,
//		layout.NewSpacer(),
//	)
//	myWindow.SetContent(vBox)
//	myWindow.ShowAndRun()
//}

func mobileLayout(username, password *widget.Entry, login *widget.Button) *fyne.Container {
	return container.New(nil)
}
func desktopLayout(username, password *widget.Entry, login *widget.Button) *fyne.Container {
	return container.NewGridWithRows(3,
		layout.NewSpacer(),
		// 列
		container.NewGridWithColumns(3,
			// 填充H和V
			layout.NewSpacer(),
			// VBox
			container.NewVBox(
				username,
				password,
				login,
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}
func main() {
	a := app.New()
	w := a.NewWindow("Layouts")

	username := &widget.Entry{PlaceHolder: "username"}
	password := &widget.Entry{PlaceHolder: "password"}
	login := &widget.Button{Text: "login", Icon: theme.ConfirmIcon()}
	if a.Driver().Device().IsMobile() {
		w.SetContent(mobileLayout(username, password, login))
	} else {
		w.SetContent(desktopLayout(username, password, login))
	}
	w.ShowAndRun()
}
