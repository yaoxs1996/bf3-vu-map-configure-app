package main

import (
	"bf3/src/data"
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	maps := new(data.Maps)
	maps = data.Read("./maps.json")

	fmt.Println(maps.Maps[0].SupportMode)

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome: ")
		}),
	))

	w.ShowAndRun()
}
