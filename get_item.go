package dynamock

import (
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *GetItemExpectation) ToTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *GetItemExpectation) WithKeys(keys map[string]dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

// WillReturns - method for set desired result
func (e *GetItemExpectation) WillReturns(res dynamodb.GetItemResponse) *GetItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) GetItemRequest(input *dynamodb.GetItemInput) dynamodb.GetItemRequest{
	req := dynamodb.GetItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.GetItemExpect) > 0 {
		x := e.dynaMock.GetItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				req.Error = ErrTableExpectationMismatch

				return req
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				req.Error = ErrKeyExpectationMismatch

				return req
			}
		}

		// delete first element of expectation
		e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:0], e.dynaMock.GetItemExpect[1:]...)

		req.Data = x.output.GetItemOutput

		return req
	}

	req.Error = ErrNoExpectation

	return req
}
