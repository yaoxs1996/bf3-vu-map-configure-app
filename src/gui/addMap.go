package gui

import (
	"bf3/src/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var curSelectedMapName string
var fullMapDict map[string]string = make(map[string]string)

func CreateAddContainer() *fyne.Container {
	maps := data.Read("./maps.json")
	// fullMapDict := make(map[string]string) // key: 完整名称, value: tech_name
	fullMapName := []string{}
	for key, value := range maps {
		fullMapDict[value.Name] = key
		fullMapName = append(fullMapName, value.Name)
	}

	optionalMap := widget.NewSelect(fullMapName, changedHandler)
	btnAddNewMap := widget.NewButton("Add", tapHandler)
	container := container.NewGridWithColumns(2, optionalMap, btnAddNewMap)

	return container
}

func changedHandler(value string) {
	curSelectedMapName = value
}

func tapHandler() {
	if len(curSelectedMapName) != 0 {
		UpdateMapList(curSelectedMapName)
	}
}
