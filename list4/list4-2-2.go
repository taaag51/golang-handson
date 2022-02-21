package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		container.NewHBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample application!"),
		),
	)

	w.ShowAndRun()
}
