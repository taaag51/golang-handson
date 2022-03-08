package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(100, 100)),
			widget.NewButton("One", nil),
			widget.NewButton("Two", nil),
			widget.NewButton("Three", nil),
			widget.NewButton("Four", nil),
			// layout.NewSpacer(),
			widget.NewButton("Five", nil),
			widget.NewButton("Six", nil),
			// layout.NewSpacer(),
			widget.NewButton("Seven", nil),
			widget.NewButton("Eight", nil),
			widget.NewButton("Nine", nil),
			widget.NewButton("Tem", nil),
		),
	)
	w.ShowAndRun()
}
