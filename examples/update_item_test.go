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

func TestUpdateItem(t *testing.T) {
	ID := 33
	expectKey := map[string]dynamodb.AttributeValue{
		"id": {
			N: aws.String(strconv.Itoa(ID)),
		},
	}

	expectedResult := "morty"
	result := dynamodb.UpdateItemResponse{
		UpdateItemOutput: &dynamodb.UpdateItemOutput{
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

	Mock.ExpectUpdateItem().Table("employee").WithKeys(expectKey).WillReturn(result)

	actualResult, err := UpdateNameByID(ID, expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != expectedResult {
		t.Fatal("Test Fail", actualResult, expectedResult)
	}
}
