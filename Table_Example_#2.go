package main

import (
  "fmt"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

// BRO-IT

func main() {
  a := app.New()
  w := a.NewWindow("SUBSCRIBE :)")
  w.Resize(fyne.NewSize(500, 500))

  names := []string{
    "Kate",
    "John",
    "Andrew",
    "Vlad",
  }

  marks := [][]int{
    {9, 10, 7},
    {10, 11, 8},
    {11, 12, 5},
    {12, 8, 10},
  }

  table := widget.NewTable(
    func() (int, int) {
      return len(names), len(marks[0]) + 1
    },

    func() fyne.CanvasObject {
      return widget.NewLabel("Default text")
    },

    func(tci widget.TableCellID, co fyne.CanvasObject) {
      if tci.Col == 0 {
        co.(*widget.Label).SetText(names[tci.Row])
      } else {
        co.(*widget.Label).SetText(fmt.Sprint(marks[tci.Row][tci.Col-1]))
      }
    },
  )

  w.SetContent(table)
  w.Show()
  a.Run()
}
