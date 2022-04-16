package main

import (
	"bf3/src/data"
	"bf3/src/gui"
	"fmt"
)

func main() {
	maps := data.Read("./maps.json")

	fmt.Println(maps["MP_017"].CnName)

	myWindow := gui.App()

	myWindow.ShowAndRun()
}
