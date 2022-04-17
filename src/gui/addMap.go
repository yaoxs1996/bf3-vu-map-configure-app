package gui

import (
	"bf3/src/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateAddContainer() *fyne.Container {
	maps := data.Read("./maps.json")
	fullMapDict := make(map[string]string) // key: 完整名称, value: tech_name
	fullMapName := []string{}
	for key, value := range maps {
		fullMapDict[value.Name] = key
		fullMapName = append(fullMapName, value.Name)
	}

	optionalMap := widget.NewSelect(fullMapName, func(valus string) {})
	btnAddNewMap := widget.NewButton("Add", func() {})
	container := container.NewGridWithColumns(2, optionalMap, btnAddNewMap)

	return container
}
