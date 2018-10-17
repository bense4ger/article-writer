package service

import (
	"bense4ger/article-writer/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ReadArticleFromFile reads a Json file and creates and Article model
func ReadArticleFromFile(filePath string) (*model.Article, error) {
	if filePath == "" {
		return nil, fmt.Errorf("ReadArticleFromFile: No file path")
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadArticleFromFile: %s", err.Error())
	}

	m := &model.Article{}
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, fmt.Errorf("ReadArticleFromFile: %s", err.Error())
	}

	m.SetKey()

	return m, nil
}
