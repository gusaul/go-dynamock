package dynamock

import (
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

func (e *DynaMock) ExpectBatchGetItem() *BatchGetItemExpectation {
	batchGetItemExpect := BatchGetItemExpectation{input: nil}
	e.BatchGetItemExpect = append(e.BatchGetItemExpect, batchGetItemExpect)

	return &e.BatchGetItemExpect[len(e.BatchGetItemExpect)-1]
}

func (e *DynaMock) ExpectUpdateItem() *UpdateItemExpectation {
	updateItemExpect := UpdateItemExpectation{attributeUpdates: nil, table: nil, key: nil}
	e.UpdateItemExpect = append(e.UpdateItemExpect, updateItemExpect)

	return &e.UpdateItemExpect[len(e.UpdateItemExpect)-1]
}
