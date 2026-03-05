package main

import (
	"fmt"
	"main/logic"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Pomodoro Timer")

	hello := widget.NewLabel("Timer")

	w.SetContent(container.NewVBox(
		hello,
		container.NewHBox(widget.NewButton("Pomodoro", func() {
			hello.SetText("25:00")
		}),
			widget.NewButton("Short Break", func() {
				hello.SetText("5:00")
			}),
			widget.NewButton("Long Break", func() {
				hello.SetText("30:00")
			})),
		widget.NewButton("Start", func() {
			logic.StartTimer()
		}),
	))
	fmt.Println("App is runnning!")
	w.ShowAndRun()

}
