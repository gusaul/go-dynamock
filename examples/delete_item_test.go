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

func TestDeleteItem(t *testing.T) {
	ID := 128
	expectedResult := aws.String(strconv.Itoa(ID))

	item := map[string]dynamodb.AttributeValue{
		"id": {
			N: expectedResult,
		},
	}

	result := dynamodb.DeleteItemResponse{
		DeleteItemOutput: &dynamodb.DeleteItemOutput{
			Attributes: map[string]dynamodb.AttributeValue{
				"id": {
					N: aws.String(strconv.Itoa(ID)),
				},
			},
		},
	}

	Mock.ExpectDeleteItem().Table("employee").WithKeys(item).WillReturn(result)

	actualResult, err := DeleteItemByID(ID)
	if err != nil {
		t.Fatal(err)
	}

	if aws.StringValue(actualResult) != aws.StringValue(expectedResult) {
		t.Fatalf("Fail: expected: %s, got :%s", aws.StringValue(expectedResult), aws.StringValue(actualResult))
	}
}
