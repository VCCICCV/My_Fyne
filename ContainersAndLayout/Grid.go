package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:55
// USER:21005
import (
	"fmt"
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

// 网格布局以具有固定列数的网格模式布置容器的元素。项目将填充一行，直到满足列数，之后将创建一个新行。垂直空间将在每行对象之间平均分配。
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
