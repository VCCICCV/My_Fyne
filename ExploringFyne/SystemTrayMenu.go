package main

// PROJECT_NAME:Fyne
// DATE:2022/11/10 20:09
// USER:21005
import (
	"fmt"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
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

// 在右下角
func main() {
	//a := app.New()
	//w := a.NewWindow("SystemMenu")
	//if desk, ok := a.(desktop.App); ok {
	//	m := fyne.NewMenu("MyApp",
	//		fyne.NewMenuItem("Show", func() {
	//			log.Println("Tapped show")
	//		}))
	//	desk.SetSystemTrayMenu(m)
	//}
	//w.ShowAndRun()
}
