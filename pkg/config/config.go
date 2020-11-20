package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var configuration Configuration = produceConfig()

type Configuration struct {
	HOST string
	PORT string
}

// Configuration.HOST getter
func GetConfigHOST() string {
	return configuration.HOST
}

// Configuration.PORT getter
func GetConfigPORT() string {
	return configuration.PORT
}

// Retrieve config file
func produceConfig() Configuration {
	log.Println("Load configuration")

	configuration := Configuration{}

	file, _ := os.Open("config.json")
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&configuration); err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
