package main

// PROJECT_NAME:Fyne
// DATE:2022/11/10 20:25
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
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

type myEntry struct {
	widget.Entry
}

func desktopLayout2(login *widget.Button) *fyne.Container {

	return container.NewGridWithRows(3,
		layout.NewSpacer(),
		// 列
		container.NewGridWithColumns(3,
			// 填充H和V
			layout.NewSpacer(),
			// VBox
			container.NewVBox(
				login,
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}
func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}
	log.Println("Shortcut typed:", s)
}
func main() {
	a := app.New()
	w := a.NewWindow("快捷方式")
	login := &widget.Button{Text: "login", Icon: theme.ConfirmIcon()}
	if a.Driver().Device().IsMobile() {
		fmt.Print("出错了")
	} else {
		w.SetContent(desktopLayout2(login))
		w.Resize(fyne.NewSize(600, 500))
		w.ShowAndRun()
	}
}
