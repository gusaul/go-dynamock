package examples

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// GetName - example func using GetItem method
func GetName(id string) (*string, error) {
	param := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("employee"),
	}

	req := Fake.DB.GetItemRequest(param)
	if req.Error != nil {
		return nil, req.Error
	}

	var value *string
	output, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if v, ok := output.Item["name"]; ok {
		err := dynamodbattribute.Unmarshal(&v, &value)
		if err != nil {
			return value, err
		}
	}

	return value, nil
}
