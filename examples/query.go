package examples

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func QueryByID(ID int) (*string, error) {
	param := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("#id = :id"),
		ExpressionAttributeNames: map[string]string{
			"#id": "ID",
		},
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":id": {
				N: aws.String("1"),
			},
		},
		TableName: aws.String("employee"),
		Select:    dynamodb.SelectAllAttributes,
		Limit:     aws.Int64(1),
	}

	req := Fake.DB.QueryRequest(param)
	if req.Error != nil {
		return nil, req.Error
	}

	var value *string
	output, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if v, ok := output.QueryOutput.Items[0]["name"]; ok {
		err := dynamodbattribute.Unmarshal(&v, &value)
		if err != nil {
			return aws.String(""), err
		}
	}

	return value, nil
}
