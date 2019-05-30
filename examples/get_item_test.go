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

func TestGetName(t *testing.T) {
	expectKey := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String("1"),
		},
	}

	expectedResult := aws.String("rick sanchez")
	result := dynamodb.GetItemResponse{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]dynamodb.AttributeValue{
				"name": {
					S: expectedResult,
				},
			},
		},
	}

	Mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturn(result)

	actualResult, err := GetName("1")
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != aws.StringValue(expectedResult) {
		t.Fatal("Test Fail", actualResult, *expectedResult)
	}
}
