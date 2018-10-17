package main

import (
	"bense4ger/article-writer/config"
	"bense4ger/article-writer/service"
	"fmt"
)

func main() {
	err := service.Create(config.C.AWSRegion, config.C.Table)
	if err != nil {
		fmt.Printf("Error creating service: %s", err.Error())
	}

}
