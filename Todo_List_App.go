package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// BRO-IT

type Task struct {
	Id          uint
	Title       string
	Description string
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Task Manager")
	w.Resize(fyne.NewSize(500, 600))
	w.CenterOnScreen()

	var tasks []Task
	var createContent *fyne.Container
	var tasksContent *fyne.Container
	var tasksList *widget.List

	DB, _ := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	DB.AutoMigrate(&Task{})
	DB.Find(&tasks)

	noTasksLabel := canvas.NewText("No Tasks", color.Black)

	if len(tasks) != 0 {
		noTasksLabel.Hide()
	}

	newTaskIcon, _ := fyne.LoadResourceFromPath("./icons/plus.png")
	back, _ := fyne.LoadResourceFromPath("./icons/back.png")
	save, _ := fyne.LoadResourceFromPath("./icons/save.png")
	delete, _ := fyne.LoadResourceFromPath("./icons/delete.png")
	edit, _ := fyne.LoadResourceFromPath("./icons/edit.png")

	tasksBar := container.NewHBox(
		canvas.NewText("Your Tasks", color.Black),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", newTaskIcon, func() {
			w.SetContent(createContent)
		}),
	)

	tasksList = widget.NewList(
		func() int {
			return len(tasks)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("Default")
		},

		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(tasks[lii].Title)
		},
	)

	tasksList.OnSelected = func(id widget.ListItemID) {
		detailsBar := container.NewHBox(
			canvas.NewText(
				fmt.Sprintf(
					"Details about \"%s\"",
					tasks[id].Title,
				),
				color.Black,
			),
			layout.NewSpacer(),

			widget.NewButtonWithIcon("", back, func() {
				w.SetContent(tasksContent)
				tasksList.Unselect(id)
			}),
		)

		taskTitle := widget.NewLabel(tasks[id].Title)
		taskTitle.TextStyle = fyne.TextStyle{Bold: true}

		taskDescription := widget.NewLabel(tasks[id].Description)
		taskDescription.TextStyle = fyne.TextStyle{Italic: true}
		taskDescription.Wrapping = fyne.TextWrapBreak

		buttonsBox := container.NewHBox(

			// DELETE
			widget.NewButtonWithIcon(
				"",
				delete,

				func() {
					dialog.ShowConfirm(
						"Deleting task",

						fmt.Sprintf(
							"Are you sure about deleting task \"%s\"?",
							tasks[id].Title,
						),

						func(b bool) {
							if b {
								DB.Delete(&Task{}, "Id = ?", tasks[id].Id)
								DB.Find(&tasks)

								if len(tasks) == 0 {
									noTasksLabel.Show()
								} else {
									noTasksLabel.Hide()
								}
							}

							tasksList.UnselectAll()
							w.SetContent(tasksContent)
						},

						w,
					)
				},
			),

			// EDIT
			widget.NewButtonWithIcon(
				"",
				edit,

				func() {
					editBar := container.NewHBox(
						canvas.NewText(
							fmt.Sprintf(
								"Editing \"%s\"",
								tasks[id].Title,
							),
							color.Black,
						),
						layout.NewSpacer(),

						widget.NewButtonWithIcon("", back, func() {
							w.SetContent(tasksContent)
							tasksList.Unselect(id)
						}),
					)

					editTitle := widget.NewEntry()
					editTitle.SetText(tasks[id].Title)

					editDescription := widget.NewMultiLineEntry()
					editDescription.SetText(tasks[id].Description)

					editButton := widget.NewButtonWithIcon(
						"Save Task",
						save,

						// EDIT TASK FUNCTION
						func() {
							DB.Find(
								&Task{},
								"Id = ?",
								tasks[id].Id,
							).Updates(
								Task{
									Title:       editTitle.Text,
									Description: editDescription.Text,
								},
							)

							DB.Find(&tasks)

							w.SetContent(tasksContent)
							tasksList.UnselectAll()
						},
					)

					editContent := container.NewVBox(
						editBar,
						canvas.NewLine(color.Black),

						editTitle,
						editDescription,
						editButton,
					)

					w.SetContent(editContent)
				},
			),
		)

		detailsVBox := container.NewVBox(
			detailsBar,
			canvas.NewLine(color.Black),

			taskTitle,
			taskDescription,
			buttonsBox,
		)

		w.SetContent(detailsVBox)
	}

	tasksScroll := container.NewScroll(tasksList)
	tasksScroll.SetMinSize(fyne.NewSize(500, 500))

	tasksContent = container.NewVBox(
		tasksBar,
		canvas.NewLine(color.Black),
		noTasksLabel,
		tasksScroll,
	)

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Task title...")

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetPlaceHolder("Task description...")

	saveTaskButton := widget.NewButtonWithIcon("Save Task", save, func() {
		task := Task{
			Title:       titleEntry.Text,
			Description: descriptionEntry.Text,
		}

		DB.Create(&task)
		DB.Find(&tasks)

		titleEntry.Text = ""
		titleEntry.Refresh()

		descriptionEntry.Text = ""
		descriptionEntry.Refresh()

		w.SetContent(tasksContent)

		tasksList.UnselectAll()

		if len(tasks) == 0 {
			noTasksLabel.Show()
		} else {
			noTasksLabel.Hide()
		}
	})

	createBar := container.NewHBox(
		canvas.NewText("Create new task", color.Black),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", back, func() {
			titleEntry.Text = ""
			titleEntry.Refresh()

			descriptionEntry.Text = ""
			descriptionEntry.Refresh()

			w.SetContent(tasksContent)

			tasksList.UnselectAll()
		}),
	)

	createContent = container.NewVBox(
		createBar,
		canvas.NewLine(color.Black),

		container.NewVBox(
			titleEntry,
			descriptionEntry,
			saveTaskButton,
		),
	)

	w.SetContent(tasksContent)
	w.Show()
	a.Run()
}
