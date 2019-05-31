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

func TestPutItem(t *testing.T) {
	ID := 1
	expectedResult := "pepe the frog"

	item := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String(strconv.Itoa(ID)),
		},
		"name": {
			S: aws.String(expectedResult),
		},
	}

	result := dynamodb.PutItemResponse{
		PutItemOutput: &dynamodb.PutItemOutput{
			Attributes: map[string]dynamodb.AttributeValue{
				"id": {
					N: aws.String(strconv.Itoa(ID)),
				},
				"name": {
					S: aws.String(expectedResult),
				},
			},
		},
	}

	Mock.ExpectPutItem().Table("employee").WithItems(item).WillReturn(result)

	actualResult, err := PutNameByID(ID, expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != expectedResult {
		t.Fatalf("Fail: expected: %s, got :%s", expectedResult, aws.StringValue(actualResult))
	}
}
