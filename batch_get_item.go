package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
)

func (e *BatchGetItemExpectation) WithRequest(input map[string]*dynamodb.KeysAndAttributes) *BatchGetItemExpectation {
	e.input = input
	return e
}

func (e *BatchGetItemExpectation) ExpectReturns(res dynamodb.BatchGetItemOutput) *BatchGetItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	if len(e.dynaMock.BatchGetItemExpect) > 0 {
		x := e.dynaMock.BatchGetItemExpect[0] //get first element of expectation

		if x.input != nil {
			if !reflect.DeepEqual(x.input, input.RequestItems) {
				return nil, fmt.Errorf("Expect input %+v but found input %+v", x.input, input.RequestItems)
			}
		}

		// delete first element of expectation
		e.dynaMock.BatchGetItemExpect = append(e.dynaMock.BatchGetItemExpect[:0], e.dynaMock.BatchGetItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Get Item Expectation Not Found")
}
