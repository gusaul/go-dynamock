package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WithRequest - method for set Request expectation
func (e *BatchGetItemExpectation) WithRequest(input map[string]*dynamodb.KeysAndAttributes) *BatchGetItemExpectation {
	e.input = input
	return e
}

// WillReturns - method for set desired result
func (e *BatchGetItemExpectation) WillReturns(res dynamodb.BatchGetItemOutput) *BatchGetItemExpectation {
	e.output = &res
	return e
}

// BatchGetItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	if len(e.dynaMock.BatchGetItemExpect) > 0 {
		x := e.dynaMock.BatchGetItemExpect[0] //get first element of expectation

		if x.input != nil {
			if !reflect.DeepEqual(x.input, input.RequestItems) {
				return &dynamodb.BatchGetItemOutput{}, fmt.Errorf("Expect input %+v but found input %+v", x.input, input.RequestItems)
			}
		}

		// delete first element of expectation
		e.dynaMock.BatchGetItemExpect = append(e.dynaMock.BatchGetItemExpect[:0], e.dynaMock.BatchGetItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.BatchGetItemOutput{}, fmt.Errorf("Batch Get Item Expectation Not Found")
}

// BatchGetItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opt ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	if len(e.dynaMock.BatchGetItemExpect) > 0 {
		x := e.dynaMock.BatchGetItemExpect[0] //get first element of expectation

		if x.input != nil {
			if !reflect.DeepEqual(x.input, input.RequestItems) {
				return &dynamodb.BatchGetItemOutput{}, fmt.Errorf("Expect input %+v but found input %+v", x.input, input.RequestItems)
			}
		}

		// delete first element of expectation
		e.dynaMock.BatchGetItemExpect = append(e.dynaMock.BatchGetItemExpect[:0], e.dynaMock.BatchGetItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.BatchGetItemOutput{}, fmt.Errorf("Batch Get Item With Context Expectation Not Found")
}
