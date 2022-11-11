package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:54
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

// 渐变
func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")
	// 线性
	//gradient := canvas.NewHorizontalGradient(color.RGBA{20, 20, 20, 255}, color.Transparent)
	// 径向
	gradient := canvas.NewRadialGradient(color.RGBA{20, 20, 20, 255}, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
