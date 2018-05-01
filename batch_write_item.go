package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"reflect"
)

func (e *BatchWriteItemExpectation) WithRequest(input map[string][]*dynamodb.WriteRequest) *BatchWriteItemExpectation {
	e.input = input
	return e
}

func (e *BatchWriteItemExpectation) WillReturns(res dynamodb.BatchWriteItemOutput) *BatchWriteItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	if len(e.dynaMock.BatchWriteItemExpect) > 0 {
		x := e.dynaMock.BatchWriteItemExpect[0] //get first element of expectation

		if x.input != nil {
			if !reflect.DeepEqual(x.input, input.RequestItems) {
				return nil, fmt.Errorf("Expect input %+v but found input %+v", x.input, input.RequestItems)
			}
		}

		// delete first element of expectation
		e.dynaMock.BatchWriteItemExpect = append(e.dynaMock.BatchWriteItemExpect[:0], e.dynaMock.BatchWriteItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Batch Write Item Expectation Not Found")
}
