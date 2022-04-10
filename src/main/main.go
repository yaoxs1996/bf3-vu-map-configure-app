package main

import (
	"bf3/src/data"
	"fmt"
)

// Definition of map info json
type Maps struct {
	Maps []Map `json:"maps"`
}

type Map struct {
	TechName    string `json:"tech_name"`
	Name        string `json:"name"`
	CnName      string `json:"cn_name"`
	SupportMode Mode   `json:"support_mode"`
}

type Mode struct {
	Conquest64          int `json:"ConquestLarge0"`
	Conquest            int `json:"ConquestSmall0"`
	Rush                int `json:"RushLarge0"`
	SquadDeathmatch     int `json:"SquadDeathMatch0"`
	TeamDeathmatch      int `json:"TeamDeathMatch0"`
	TeamDMCloseQuarters int `json:"TeamDeathMatchC0"`
	ConquestAssault64   int `json:"ConquestAssaultLarge0"`
	ConquestAssault     int `json:"ConquestAssaultSmall0"`
	ConquestAssaultDay2 int `json:"ConquestAssaultSmall1"`
	ConquestDomination  int `json:"Domination0"`
	GunMaster           int `json:"GunMaster0"`
	TankSuperiority     int `json:"TankSuperiority0"`
	SquadRush           int `json:"SquadRush0"`
	CaptureTheFlag      int `json:"CaptureTheFlag0"`
}

func main() {
	maps = make(Maps)

	maps = data.Read("./maps.json")

	fmt.Println(maps.Maps[0].SupportMode.CaptureTheFlag)
}
