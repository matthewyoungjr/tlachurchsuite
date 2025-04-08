package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TLA Church Suite")
	w.Resize(fyne.Size{Width: 900, Height: 650})

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()

}
