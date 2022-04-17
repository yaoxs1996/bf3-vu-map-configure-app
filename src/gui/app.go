package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func App() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("BF3 VU Map")

	// layout := layout.NewVBoxLayout()

	// 添加新地图控件
	addContainer := new(fyne.Container)
	addContainer = CreateAddContainer()

	// addCard := widget.NewCard("Add new map", "", addContainer)

	// 展示控件
	list := new(widget.List)
	list = ListMap()

	// showCard := widget.NewCard("Map and mode", "Current", list)

	// 装配
	btn := widget.NewButton("123", func() {})
	myWindow.SetContent(container.NewGridWithRows(2, list, container.NewVBox(addContainer, btn)))

	return myWindow
}
