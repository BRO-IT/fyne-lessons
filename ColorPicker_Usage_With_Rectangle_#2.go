package main

import (
  "image/color"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/widget"
)
// BRO-IT
func main() {
  a := app.NewWithID("BRO-IT APP")
  w := a.NewWindow("Подпишись на канал :)")
  w.Resize(fyne.NewSize(700, 700))

  rec := canvas.NewRectangle(color.White)
  rec.SetMinSize(fyne.NewSize(300, 300))

  cp := dialog.NewColorPicker(
    "Color Picker",
    "Chhose the color",
    func(c color.Color) {
      rec.FillColor = c
      rec.Refresh()
    },
    w,
  )

  btn := widget.NewButton("Call ColorPicker", func() {
    cp.Show()
  })

  w.SetContent(
    container.NewVBox(
      btn,
      rec,
    ),
  )
  w.Show()
  a.Run()
}
