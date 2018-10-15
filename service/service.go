package service

import (
	"bense4ger/article-writer/model"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	svc   *dynamodb.DynamoDB
	table string
)

// Create instatiates the DynamoDB service
func Create(region, tbl string) error {
	if table == "" {
		return fmt.Errorf("Create: No table name supplied")
	}
	table = tbl

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return fmt.Errorf("Create: %s", err.Error())
	}

	svc = dynamodb.New(sess)
	return nil
}

// PutArticle writes and article model to Dynamo
func PutArticle(article *model.Article) error {
	av, err := dynamodbattribute.MarshalMap(article)
	if err != nil {
		return fmt.Errorf("PutArticle: %s", err.Error())
	}

	pii := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = svc.PutItem(pii)
	if err != nil {
		return fmt.Errorf("PutArticle: %s", err.Error())
	}

	return nil
}
