package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// WithRequest - method for set Request expectation
func (e *BatchWriteItemExpectation) WithRequest(input map[string][]dynamodb.WriteRequest) *BatchWriteItemExpectation {
	e.input = input
	return e
}

// WillReturns - method for set desired result
func (e *BatchWriteItemExpectation) WillReturn(res dynamodb.BatchWriteItemResponse) *BatchWriteItemExpectation {
	e.output = res.BatchWriteItemOutput
	return e
}

// BatchWriteItemRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) dynamodb.BatchWriteItemRequest {
	req := dynamodb.BatchWriteItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.BatchWriteItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.BatchWriteItemExpect[0]

	validateInput(input, req.Request)
	validateItem(x.input, input.RequestItems, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.BatchWriteItemExpect = append(e.dynaMock.BatchWriteItemExpect[:0], e.dynaMock.BatchWriteItemExpect[1:]...)

	req.Data = x.output

	return req
}
