package examples

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamock "go-dynamock"
)

var mock *dynamock.DynaMock

func init() {
	Fake = new(FakeDynamo)
	Fake.DB, mock = dynamock.New()
}

func TestGetName(t *testing.T) {
	expectKey := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String("1"),
		},
	}

	expectedResult := aws.String("jaka")
	result := dynamodb.GetItemResponse{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]dynamodb.AttributeValue{
				"name": {
					S: expectedResult,
				},
			},
		},
	}

	//lets start dynamock in action
	mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturns(result)

	actualResult, err := GetName("1")
	if err != nil{
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != aws.StringValue(expectedResult) {
		t.Fatal("Test Fail", actualResult, *expectedResult)
	}
}
