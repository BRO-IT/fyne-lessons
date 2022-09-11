package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/widget"
)

// BRO-IT

func main() {
  a := app.New()
  w := a.NewWindow("Подпишись на канал :)")
  w.Resize(fyne.NewSize(650, 650))

  title := widget.NewLabel("Sign Up Progress")

  pb := widget.NewProgressBar()
  pb.Max = 100.0

  name := widget.NewEntry()
  name.SetPlaceHolder("Your name")

  surname := widget.NewEntry()
  surname.SetPlaceHolder("Your surname")

  phone_number := widget.NewEntry()
  phone_number.SetPlaceHolder("Your phone number")

  nick := widget.NewEntry()
  nick.SetPlaceHolder("Your nick")

  password := widget.NewPasswordEntry()
  password.SetPlaceHolder("Your password")

  email := widget.NewEntry()
  email.SetPlaceHolder("Your email")

  bio := widget.NewMultiLineEntry()
  bio.SetPlaceHolder("Info about you")

  tabs := container.NewAppTabs(
    container.NewTabItem(
      "Personal Info",
      container.NewVBox(
        name,
        surname,
        phone_number,
      ),
    ),
    container.NewTabItem(
      "Data for Sign Up",
      container.NewVBox(
        nick,
        password,
      ),
    ),
    container.NewTabItem(
      "Additional Info",
      container.NewVBox(
        email,
        bio,
      ),
    ),
  )

  next := widget.NewButton("Next", func() {
    if tabs.SelectedIndex() < 2 {
      tabs.SelectIndex(tabs.SelectedIndex() + 1)
      pb.SetValue(pb.Value + 33)
    } else {
      pb.SetValue(pb.Value + 34)

      dialog.ShowInformation(
        "Sign Up",
        "You've beeen signed up!",
        w,
      )
    }
  })

  w.SetContent(
    container.NewVBox(
      title,
      pb,
      tabs,
      next,
    ),
  )
  w.Show()
  a.Run()
}
