package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

/*
1) Знакомство с ProgressBar
2) Изменение значения(состояния) ProgressBar
3) Получение значения ProgressBar
4) Практическое применение ProgressBar
5) Виджет ProgressBarInfinite
6) Управление анимацией ProgressBarInfinite
7) Практическое применение ProgressBarInfinite
*/

func main() {
	a := app.New()
	w := a.NewWindow("Подпишись на канал :)")
	w.Resize(fyne.NewSize(650, 650))

	pbinf := widget.NewProgressBarInfinite()
	pbinf.Hide()

	title := widget.NewLabel("Create your post")

	post_title := widget.NewEntry()
	post_title.SetPlaceHolder("Your post title")

	post_text := widget.NewMultiLineEntry()
	post_text.SetPlaceHolder("Your post text")

	submit := widget.NewButton("Submit", func() {
		pbinf.Show()
		time.Sleep(time.Second * 3)
		pbinf.Hide()

		dialog.ShowInformation(
			"Post Creation",
			"You have created your post!",
			w,
		)
	})

	w.SetContent(
		container.NewVBox(
			pbinf,
			title,
			post_title,
			post_text,
			submit,
		),
	)
	w.Show()
	a.Run()
}
