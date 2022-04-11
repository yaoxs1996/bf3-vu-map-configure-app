package main

import (
	"bf3/src/data"
	"fmt"
)

func main() {
	maps := new(data.Maps)
	maps = data.Read("./maps.json")

	fmt.Println(maps.Maps[0].SupportMode)
}
