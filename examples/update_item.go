package examples

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// UpdateName - example func using UpdateItemRequest method
func UpdateNameByID(ID int, name string) (*string, error) {
	param := &dynamodb.UpdateItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {N: aws.String(strconv.Itoa(ID))},
		},
		AttributeUpdates: map[string]dynamodb.AttributeValueUpdate{
			"name": {Action: dynamodb.AttributeActionPut, Value: &dynamodb.AttributeValue{N: aws.String(name)}},
		},
		TableName: aws.String("employee"),
	}

	req := Fake.DB.UpdateItemRequest(param)
	if req.Error != nil {
		return nil, req.Error
	}

	var value *string
	output, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if v, ok := output.UpdateItemOutput.Attributes["name"]; ok {
		err := dynamodbattribute.Unmarshal(&v, &value)
		if err != nil {
			return aws.String(""), err
		}
	}

	return value, nil
}
