package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:54
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"image/color"
	"os"
	"time"
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
	w := a.NewWindow("Hello")
	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}
	green := color.NRGBA{G: 0xff, A: 0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*3, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()
	canvas.NewColorRGBAAnimation(blue, green, time.Second*3, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()
	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false)
	w.ShowAndRun()
}
