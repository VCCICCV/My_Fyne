package main

// PROJECT_NAME:Fyne
// DATE:2022/11/11 18:34
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
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
	w.Resize(fyne.NewSize(600, 500))
	str := binding.NewString()
	go func() {
		dots := ".........."
		for i := 10; i >= 0; i-- {
			str.Set("倒计时：" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}
