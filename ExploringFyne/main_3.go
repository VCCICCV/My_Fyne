//package main
//
//import (
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/widget"
//	"github.com/flopp/go-findfont"
//	"os"
//	"strings"
//)
//
//func init() {
//	fontPaths := findfont.List()
//	for _, path := range fontPaths {
//		//fmt.Println(path)
//		//楷体:simkai.ttf
//		//黑体:simhei.ttf
//		if strings.Contains(path, "simkai.ttf") {
//			fmt.Println(path)
//			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
//			break
//		}
//	}
//	fmt.Println("=============")
//}
//
//func main() {
//	MyApp := app.New()
//	c := MyApp.NewWindow("解决fyne支持中文")
//	labels := widget.NewLabel("支持中文")
//	c.SetContent(labels)
//	c.Resize(fyne.NewSize(600, 600))
//	c.ShowAndRun()
//}

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
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
func main() {
	// use the font...
	MyApp := app.New()
	c := MyApp.NewWindow("fyne支持中文")
	labels := widget.NewLabel("中文看看行不行")
	c.SetContent(labels)
	c.Resize(fyne.NewSize(600, 600))
	c.ShowAndRun()
}
