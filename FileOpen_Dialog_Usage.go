package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// BRO-IT

func main() {
	a := app.New()
	w := a.NewWindow("App")
	w.Resize(fyne.NewSize(800, 500))

	entry := widget.NewMultiLineEntry()
	entry.Resize(fyne.NewSize(600, 300))
	entry.Move(fyne.NewPos(100, 135))

	btn := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(
			func(r fyne.URIReadCloser, err error) {
				data, _ := io.ReadAll(r)
				entry.SetText(string(data))
			},
			w,
		)
	})
	btn.Resize(fyne.NewSize(150, 75))
	btn.Move(fyne.NewPos(325, 30))

	cont := container.NewWithoutLayout(
		btn,
		entry,
	)

	w.SetContent(cont)
	w.Show()
	a.Run()
}
