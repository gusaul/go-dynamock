package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WillReturns - method for set desired result
func (e *TransactWriteItemsExpectation) WillReturns(res dynamodb.TransactWriteItemsOutput) *TransactWriteItemsExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error){
	if len(e.dynaMock.TransactWriteItemsExpect) > 0 {
		x := e.dynaMock.TransactWriteItemsExpect[0] //get first element of expectation

		// delete first element of expectation
		e.dynaMock.TransactWriteItemsExpect = append(e.dynaMock.TransactWriteItemsExpect[:0],
			e.dynaMock.TransactWriteItemsExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Transact Write Items Table Expectation Not Found")
}