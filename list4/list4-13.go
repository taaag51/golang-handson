package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		container.NewVBox(
			l,
			// container.NewAppTabs(
			container.NewDocTabs(
				container.NewTabItem("First",
					widget.NewLabel("This is First tab item."),
				),
				container.NewTabItem("Second",
					widget.NewLabel("This is Second tab item."),
				),
			),
		),
	)
	w.ShowAndRun()
}
