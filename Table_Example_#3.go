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

  data := map[string][]int{
    "Denis": {10, 9, 12, 10, 8},
    "Ivan":  {9, 8, 7, 10, 11},
    "Vlad":  {10, 7, 5, 6, 12},
    "Name":  {3, 2, 1, 4, 10},
  }

  var kx string
  var names []string
  for k := range data {
    names = append(names, k)
    kx = k
  }

  table := widget.NewTable(
    func() (int, int) {
      return len(names), len(data[kx]) + 1
    },

    func() fyne.CanvasObject {
      return widget.NewLabel("Default text")
    },

    func(tci widget.TableCellID, co fyne.CanvasObject) {
      if tci.Col == 0 {
        co.(*widget.Label).SetText(names[tci.Row])
      } else {
        co.(*widget.Label).SetText(fmt.Sprint(data[names[tci.Row]][tci.Col-1]))
      }
    },
  )

  w.SetContent(table)
  w.Show()
  a.Run()
}
