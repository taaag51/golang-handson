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
	ck := widget.NewCheckGroup([]string{"check1", "check2"}, nil)
	w.SetContent(
		container.NewVBox(
			l, ck,
			widget.NewButton("OK", func() {
				re := "result: "
				if ck1.Checked {
					re += "Check-1"
				}
				if ck2.Checked {
					re += "Check-2"
				}

				l.SetText(re)
			}),
		),
	)
	w.ShowAndRun()
}
