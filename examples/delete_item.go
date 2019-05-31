package examples

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// DeleteItemByID - example func using DeleteItemRequest method
func DeleteItemByID(ID int) (*string, error) {
	param := &dynamodb.DeleteItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(ID)),
			},
		},
		TableName: aws.String("employee"),
	}

	req := Fake.DB.DeleteItemRequest(param)
	if req.Error != nil {
		return nil, req.Error
	}

	output, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	var value *string
	if v, ok := output.DeleteItemOutput.Attributes["id"]; ok {
		err := dynamodbattribute.Unmarshal(&v, &value)
		if err != nil {
			return nil, err
		}
	}

	return value, nil
}
