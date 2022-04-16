package gui

import (
	"bf3/src/data"
	"bufio"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var t = []string{"a", "string", "list"}

func ListMap() *widget.List {
	f, err := os.Open("D:\\PC\\Documents\\Battlefield 3\\Server\\Admin\\MapList.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	maps := data.Read("./maps.json")

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), " ")
		fmt.Println("map name: " + maps[info[0]].CnName + " gamemode: " + info[1] + " round: " + info[2])
		// fmt.Println(len(info))
	}

	list := widget.NewList(
		func() int {
			return len(t)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(t[i])
		})

	return list
}
