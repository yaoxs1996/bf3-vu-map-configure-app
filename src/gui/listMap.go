package gui

import (
	"bf3/src/data"
	"bufio"
	"fmt"
	"os"
	"reflect"
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
		mapInfo["techName"] = string(info[0])
		mapInfo["mode"] = string(info[1])
		mapInfo["round"] = string(info[2])
		curMapList = append(curMapList, mapInfo)
		// fmt.Println("map name: " + maps[info[0]].Name + " gamemode: " + info[1] + " round: " + info[2])
		// fmt.Println(len(info))
		nlines += 1
	}

	// fmt.Println(curMapList)

	supportModes := make(map[string][]string)

	for key, value := range maps {
		modes := value.SupportMode
		modeValue := reflect.ValueOf(modes)   // 反射获得的模式支持数值
		modeFullName := reflect.TypeOf(modes) // 反射获得的模式名称

		modeList := []string{}
		for i := 0; i < modeValue.NumField(); i++ {
			// fmt.Println(modeFullName.Field(i).Name) // 模式正式名称
			// fmt.Println(modeValue.Field(i))         // 模式对应数值

			if modeValue.Field(i).Int() > 0 {
				curName := modeFullName.Field(i).Name
				// 获得json字段对应名称
				if techName, ok := modeFullName.FieldByName(curName); ok {
					modeList = append(modeList, techName.Tag.Get("json"))
				}
			}
			// modeName := nmodes.Field(i).Name
			// if txtType, ok := nmodes.FieldByName(modeName); ok {
			// 	fmt.Println(txtType.Tag.Get("json"))
			// }
		}
		supportModes[key] = modeList
	}

	// fmt.Println(supportModes)

	list := widget.NewList(
		func() int {
			return nlines
		},
		func() fyne.CanvasObject {
			mapName := widget.NewLabel("blank")
			mode := widget.NewSelect([]string{}, func(value string) {})
			round := widget.NewEntry()

			content := container.New(layout.NewGridLayout(3), mapName, mode, round)
			return content
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(curMapList[i]["name"])
			// o.(*fyne.Container).Objects[1].(*widget.SelectEntry).SetOptions(supportModes[curMapList[i]["techName"]])
			o.(*fyne.Container).Objects[1].(*widget.Select).Options = supportModes[curMapList[i]["techName"]]
			o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(curMapList[i]["mode"])
			o.(*fyne.Container).Objects[2].(*widget.Entry).SetText(curMapList[i]["round"])
		})

	return list
}
