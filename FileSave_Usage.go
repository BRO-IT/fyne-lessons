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
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(700, 700))

	entry := widget.NewMultiLineEntry()

	save_file := widget.NewButton("Save File", func() {
		dialog.ShowFileSave(
			func(uc fyne.URIWriteCloser, err error) {
				io.WriteString(uc, entry.Text)
			},
			w,
		)
	})

	w.SetContent(
		container.NewVBox(
			save_file,
			entry,
		),
	)
	w.Show()
	a.Run()
}
