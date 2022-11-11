package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:55
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"image/color"
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
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	text4 := canvas.NewText("4", color.Black)
	text5 := canvas.NewText("5", color.Black)
	text6 := canvas.NewText("6", color.Black)
	// 设置网格宽高，自动调整
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(5, 50)),
		text1, text2, text3, text4, text5, text6)
	myWindow.SetContent(grid)

	myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
