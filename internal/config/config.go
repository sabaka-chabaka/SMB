package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SMBAddress string `json:"smb_address"`
	SMBPort    int    `json:"smb_port"`
	SMBUser    string `json:"smb_user"`
	SMBPass    string `json:"smb_pass"`
}

var Configuration Config

func Load() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("config.json not found")
		}
		panic(err)
	}

	jsonErr := json.Unmarshal(file, &Configuration)
	if jsonErr != nil {
		panic(jsonErr)
	}
}
