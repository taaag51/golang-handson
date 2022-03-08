package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		container.NewVScroll(
			container.NewVBox(
				widget.NewButton("One", nil),
				widget.NewButton("Two", nil),
				widget.NewButton("Three", nil),
				widget.NewButton("Four", nil),
				widget.NewButton("Five", nil),
				widget.NewButton("Six", nil),
				widget.NewButton("Seven", nil),
				widget.NewButton("Eight", nil),
				widget.NewButton("Nine", nil),
				widget.NewButton("Ten", nil),
			),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
