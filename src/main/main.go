package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("./maps.json")
	if err != nil {
		fmt.Println("error opening json file", err)
		return
	}
	defer jsonFile.Close()
}
