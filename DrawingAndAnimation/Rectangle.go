package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:53
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"image/color"
	"log"
	"os"
)

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
func main() {
	a := app.New()
	w := a.NewWindow("Rectangle")
	// 桌面下标
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				log.Println("Tapped show")
			}))
		desk.SetSystemTrayMenu(m)
	}
	rect1 := canvas.NewRectangle(color.Black)
	rect1.Resize(fyne.NewSize(50, 50))
	rect2 := canvas.NewRectangle(color.Black)
	vBox := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		rect1,
		layout.NewSpacer(),
		rect2,
		layout.NewSpacer(),
	)
	w.SetContent(vBox)
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}
