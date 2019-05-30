package examples

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// FakeDynamo struct hold dynamodb connection
type FakeDynamo struct {
	DB dynamodbiface.ClientAPI
}

// Fake - object from MyDynamo
var Fake *FakeDynamo

// GetName - example func using GetItem method
func GetName(id string) (*string, error) {
	parameter := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("employee"),
	}

	req := Fake.DB.GetItemRequest(parameter)
	if req.Error != nil {
		return aws.String(""), req.Error
	}

	value := aws.String("")
	if output, err := req.Send(context.Background()); err == nil {
		if v, ok := output.Item["name"]; ok {
			err := dynamodbattribute.Unmarshal(&v, &value)
			if err != nil{
				return aws.String(""), err
			}
		}
	}

	return value, nil
}
