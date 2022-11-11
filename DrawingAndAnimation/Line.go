package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:53
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.Black)
	line.Move(fyne.NewPos(20, 20))
	line.StrokeWidth = 10
	w.SetContent(line)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
