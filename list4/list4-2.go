package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample application!"),
		),
	)

	w.ShowAndRun()
}
