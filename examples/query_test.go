package examples

import (
	"testing"

	dynamock "go-dynamock"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func init() {
	Fake = new(FakeDynamo)
	Fake.DB, Mock = dynamock.New()
}

func TestQueryByID(t *testing.T) {
	expectedResult := aws.String("pepe")

	result := dynamodb.QueryResponse{
		QueryOutput: &dynamodb.QueryOutput{
			Items: []map[string]dynamodb.AttributeValue{
				{
					"name": {
						S: expectedResult,
					},
				},
			},
		},
	}

	Mock.ExpectQuery().Table("employee").WillReturn(result)

	actualResult, err := QueryByID(1)
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != aws.StringValue(expectedResult) {
		t.Fatal("Test Fail", actualResult, *expectedResult)
	}
}
