package gui

import (
	"bf3/src/data"
	"bufio"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var t = []string{"a", "string", "list"}

func ListMap() *widget.List {
	// f, err := os.Open("D:\\PC\\Documents\\Battlefield 3\\Server\\Admin\\MapList.txt")
	f, err := os.Open("C:\\Users\\Infin\\Documents\\Battlefield 3\\Server\\Admin\\MapList.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	maps := data.Read("./maps.json")

	var curMapList []map[string]string
	var nlines int = 0

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), " ")
		mapInfo := make(map[string]string)
		mapInfo["name"] = maps[info[0]].Name
		mapInfo["mode"] = string(info[1])
		mapInfo["round"] = string(info[2])
		curMapList = append(curMapList, mapInfo)
		fmt.Println("map name: " + maps[info[0]].Name + " gamemode: " + info[1] + " round: " + info[2])
		// fmt.Println(len(info))
		nlines += 1
	}

	fmt.Println(curMapList)

	list := widget.NewList(
		func() int {
			return nlines
		},
		func() fyne.CanvasObject {
			mapName := widget.NewLabel("blank")
			mode := widget.NewSelectEntry([]string{})
			round := widget.NewEntry()

			// mapName.Resize(fyne.NewSize(200, 20))
			// mode.Resize(fyne.NewSize(50, 20))
			// round.Resize(fyne.NewSize(50, 20))

			content := container.New(layout.NewGridLayout(3), mapName, mode, round)
			return content
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(curMapList[i]["name"])
			var options []string = []string{"TDM", "2"}
			o.(*fyne.Container).Objects[1].(*widget.SelectEntry).SetOptions(options)
			o.(*fyne.Container).Objects[2].(*widget.Entry).SetText("round")
		})

	return list
}
