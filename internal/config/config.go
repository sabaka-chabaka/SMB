package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SMBPort string `json:"smb_port"`
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
