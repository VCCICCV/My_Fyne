package main

// PROJECT_NAME:Fyne
// DATE:2022/11/10 9:56
// USER:21005
import (
	"fmt"
	"fyne.io/fyne/v2/app"
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
	a := app.NewWithID("com.example.tutorial.preferences")
	w := a.NewWindow("Timeout")

	var timeout time.Duration

	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		a.Preferences().SetString("AppTimeout", selected)
	})

	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

	go func() {
		time.Sleep(timeout)
		a.Quit()
	}()

	w.SetContent(timeoutSelector)
	w.ShowAndRun()
}
