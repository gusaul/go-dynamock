package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var mock *MockDynamoDB

// New - constructor for mock instantiation
// Return : 1st => DynamoDBAPI implementation, used to inject app object
// 			2nd => mock object, used to set expectation and desired result
func New() (dynamodbiface.DynamoDBAPI, *DynaMock) {
	mock = new(MockDynamoDB)
	mock.dynaMock = new(DynaMock)

	return mock, mock.dynaMock
}

// ExpectGetItem - method to start do expectation
func (e *DynaMock) ExpectGetItem() *GetItemExpectation {
	getItemExpect := GetItemExpectation{table: nil, key: nil}
	e.GetItemExpect = append(e.GetItemExpect, getItemExpect)

	return &e.GetItemExpect[len(e.GetItemExpect)-1]
}

// ExpectBatchGetItem - method to start do expectation
func (e *DynaMock) ExpectBatchGetItem() *BatchGetItemExpectation {
	batchGetItemExpect := BatchGetItemExpectation{input: nil}
	e.BatchGetItemExpect = append(e.BatchGetItemExpect, batchGetItemExpect)

	return &e.BatchGetItemExpect[len(e.BatchGetItemExpect)-1]
}

// ExpectUpdateItem - method to start do expectation
func (e *DynaMock) ExpectUpdateItem() *UpdateItemExpectation {
	updateItemExpect := UpdateItemExpectation{updateExpression: nil, table: nil, key: nil}
	e.UpdateItemExpect = append(e.UpdateItemExpect, updateItemExpect)

	return &e.UpdateItemExpect[len(e.UpdateItemExpect)-1]
}

// ExpectPutItem - method to start do expectation
func (e *DynaMock) ExpectPutItem() *PutItemExpectation {
	putItemExpect := PutItemExpectation{table: nil, item: nil}
	e.PutItemExpect = append(e.PutItemExpect, putItemExpect)

	return &e.PutItemExpect[len(e.PutItemExpect)-1]
}

// ExpectDeleteItem - method to start do expectation
func (e *DynaMock) ExpectDeleteItem() *DeleteItemExpectation {
	deleteItemExpect := DeleteItemExpectation{table: nil, key: nil}
	e.DeleteItemExpect = append(e.DeleteItemExpect, deleteItemExpect)

	return &e.DeleteItemExpect[len(e.DeleteItemExpect)-1]
}

// ExpectBatchWriteItem - method to start do expectation
func (e *DynaMock) ExpectBatchWriteItem() *BatchWriteItemExpectation {
	batchWriteItemExpect := BatchWriteItemExpectation{input: nil}
	e.BatchWriteItemExpect = append(e.BatchWriteItemExpect, batchWriteItemExpect)

	return &e.BatchWriteItemExpect[len(e.BatchWriteItemExpect)-1]
}

// ExpectCreateTable - method to start do expectation
func (e *DynaMock) ExpectCreateTable() *CreateTableExpectation {
	createTableExpect := CreateTableExpectation{keySchema: nil, table: nil}
	e.CreateTableExpect = append(e.CreateTableExpect, createTableExpect)

	return &e.CreateTableExpect[len(e.CreateTableExpect)-1]
}

// ExpectDescribeTable - method to start do expectation
func (e *DynaMock) ExpectDescribeTable() *DescribeTableExpectation {
	describeTableExpect := DescribeTableExpectation{table: nil}
	e.DescribeTableExpect = append(e.DescribeTableExpect, describeTableExpect)

	return &e.DescribeTableExpect[len(e.DescribeTableExpect)-1]
}

// ExpectWaitTableExist - method to start do expectation
func (e *DynaMock) ExpectWaitTableExist() *WaitTableExistExpectation {
	waitTableExistExpect := WaitTableExistExpectation{table: nil}
	e.WaitTableExistExpect = append(e.WaitTableExistExpect, waitTableExistExpect)

	return &e.WaitTableExistExpect[len(e.WaitTableExistExpect)-1]
}

// ExpectScan - method to start do expectation
func (e *DynaMock) ExpectScan() *ScanExpectation {
	ScanExpect := ScanExpectation{table: nil}
	e.ScanExpect = append(e.ScanExpect, ScanExpect)

	return &e.ScanExpect[len(e.ScanExpect)-1]
}

// ExpectQuery - method to start do expectation
func (e *DynaMock) ExpectQuery() *QueryExpectation {
	queryExpect := QueryExpectation{table: nil}
	e.QueryExpect = append(e.QueryExpect, queryExpect)

	return &e.QueryExpect[len(e.QueryExpect)-1]
}

// ExpectTransactWriteItems - method to start do expectation
func (e *DynaMock) ExpectTransactWriteItems() *TransactWriteItemsExpectation {
	transactWriteItemsExpect := TransactWriteItemsExpectation{table: nil}
	e.TransactWriteItemsExpect = append(e.TransactWriteItemsExpect, transactWriteItemsExpect)

	return &e.TransactWriteItemsExpect[len(e.TransactWriteItemsExpect)-1]
}
