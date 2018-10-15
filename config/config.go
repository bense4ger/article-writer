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
		fmt.Printf("Error reading config: %s", err.Error())
		os.Exit(1)
	}

	C = &Config{}
	json.Unmarshal(b, C)
}
