package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

// BRO-IT 

func main() {
  a := app.New()
  w := a.NewWindow("SUBSCRIBE :)")
  w.Resize(fyne.NewSize(500, 500))

  data := [][]string{
    {"John", "12", "12", "12"},
    {"Andrew", "11", "11", "11"},
    {"Kate", "10", "10", "10"},
  }

  table := widget.NewTable(
    func() (int, int) {
      return len(data), len(data[0])
    },

    func() fyne.CanvasObject {
      return widget.NewLabel("Default text")
    },

    func(tci widget.TableCellID, co fyne.CanvasObject) {
      co.(*widget.Label).SetText(data[tci.Row][tci.Col])
    },
  )

  w.SetContent(table)
  w.Show()
  a.Run()
}
