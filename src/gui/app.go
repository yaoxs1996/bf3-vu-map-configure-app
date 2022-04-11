package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func App() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("BF3 VU Map")

	list := new(widget.List)
	list = ListMap()

	myWindow.SetContent(list)

	return myWindow
}
