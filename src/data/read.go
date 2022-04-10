package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func Read(filePath string) *Maps {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Successfully read maps.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var maps Maps

	json.Unmarshal(byteValue, &maps)

	return &maps
}
