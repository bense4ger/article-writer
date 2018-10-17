package main

import (
	"bense4ger/article-writer/config"
	"bense4ger/article-writer/service"
	"flag"
	"fmt"
	"os"
)

var (
	filePath string
)

func main() {
	err := service.Create(config.C.AWSRegion, config.C.Table)
	if err != nil {
		fmt.Printf("Error creating service: %s\n", err.Error())
		os.Exit(1)
	}

	article, err := service.ReadArticleFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading article from file: %s\n", err.Error())
		os.Exit(1)
	}

	err = service.PutArticle(article)
	if err != nil {
		fmt.Printf("Error writing to Dyanmo: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Article published")
	os.Exit(0)
}

func init() {
	flag.StringVar(&filePath, "filePath", "", "The file path of the article to write to Dynamo")
	flag.Parse()
}
