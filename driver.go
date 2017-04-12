package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var mock *MockDynamoDB

func New() (dynamodbiface.DynamoDBAPI, *DynaMock) {
	mock = new(MockDynamoDB)
	mock.dynaMock = new(DynaMock)

	return mock, mock.dynaMock
}

func (e *DynaMock) ExpectGetItem() *GetItemExpectation {
	getItemExpect := GetItemExpectation{table: nil, key: nil}
	e.GetItemExpect = append(e.GetItemExpect, getItemExpect)

	return &e.GetItemExpect[len(e.GetItemExpect)-1]
}
