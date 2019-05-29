package examples

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

// MyDynamo struct hold dynamodb connection
type MyDynamo struct {
	Db dynamodbiface.ClientAPI
}

// Dyna - object from MyDynamo
var Dyna *MyDynamo

// ConfigureDynamoDB - init func for open connection to aws dynamodb
func ConfigureDynamoDB() {
	Dyna = new(MyDynamo)
	cnf, _ := external.LoadDefaultAWSConfig()
	Dyna.Db = dynamodb.New(cnf)
}

// GetName - example func using GetItem method
func GetName(ctx context.Context, id string) (*dynamodb.GetItemResponse, error) {
	parameter := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("employee"),
	}

	req := Dyna.Db.GetItemRequest(parameter)
	response, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	return response, nil
}
