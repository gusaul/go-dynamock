package examples

import (
	"strconv"
	"testing"

	dynamock "github.com/groovili/go-dynamock"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func init() {
	Fake = new(FakeDynamo)
	Fake.DB, Mock = dynamock.New()
}

func TestGetItem(t *testing.T) {
	ID := 123
	expectKey := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String(strconv.Itoa(ID)),
		},
	}

	expectedResult := "rick sanchez"
	result := dynamodb.GetItemResponse{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]dynamodb.AttributeValue{
				"id": {
					N: aws.String(strconv.Itoa(ID)),
				},
				"name": {
					S: aws.String(expectedResult),
				},
			},
		},
	}

	Mock.ExpectGetItem().Table("employee").WithKeys(expectKey).WillReturn(result)

	actualResult, err := GetNameByID(ID)
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != expectedResult {
		t.Fatalf("Fail: expected: %s, got: %s", expectedResult, aws.StringValue(actualResult))
	}
}
