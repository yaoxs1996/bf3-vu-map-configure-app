package main

import (
	"bf3/src/data"
	"bf3/src/gui"
	"fmt"

	"fyne.io/fyne/v2"
)

func main() {
	maps := data.Read("./maps.json")

	fmt.Println(maps["MP_017"].CnName)

	myWindow := gui.App()
	myWindow.Resize(fyne.NewSize(600, 800))

	myWindow.ShowAndRun()
}
