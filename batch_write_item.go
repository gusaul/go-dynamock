package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithRequest - method for set Request expectation
func (e *BatchWriteItemExpectation) WithRequest(input map[string][]*dynamodb.WriteRequest) *BatchWriteItemExpectation {
	e.input = input
	return e
}

// WillReturns - method for set desired result
func (e *BatchWriteItemExpectation) WillReturns(res dynamodb.BatchWriteItemOutput) *BatchWriteItemExpectation {
	e.output = &res
	return e
}

// BatchWriteItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	if len(e.dynaMock.BatchWriteItemExpect) > 0 {
		for i, x := range e.dynaMock.BatchWriteItemExpect {
			if x.input != nil {
				if reflect.DeepEqual(x.input, input.RequestItems) {
					e.dynaMock.BatchWriteItemExpect = append(e.dynaMock.BatchWriteItemExpect[:i], e.dynaMock.BatchWriteItemExpect[i:]...)
					return x.output, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Batch Write Item Expectation Failed. Expected one of %+v to equal %+v", e.dynaMock.BatchWriteItemExpect, input.RequestItems)
}

// BatchWriteItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opt ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
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
