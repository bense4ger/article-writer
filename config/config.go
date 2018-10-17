package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// C is an exported config, configured on package init
var C *Config

// Config models configuration data
type Config struct {
	AWSRegion string
	Table     string
}

func init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Error reading config: %s\n", err.Error())
		os.Exit(1)
	}

	C = &Config{}
	err = json.Unmarshal(b, C)
	if err != nil {
		fmt.Printf("Error unmarshalling config: %s\n", err.Error())
		os.Exit(1)
	}
}
