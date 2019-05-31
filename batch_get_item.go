package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// WithRequest - method for set Request expectation
func (e *BatchGetItemExpectation) WithRequest(input map[string]dynamodb.KeysAndAttributes) *BatchGetItemExpectation {
	e.input = input
	return e
}

// WillReturn - method for set desired result
func (e *BatchGetItemExpectation) WillReturn(res dynamodb.BatchGetItemResponse) *BatchGetItemExpectation {
	e.output = res.BatchGetItemOutput
	return e
}

// BatchGetItemRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) dynamodb.BatchGetItemRequest {
	req := dynamodb.BatchGetItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.BatchGetItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.BatchGetItemExpect[0]

	validateInput(input, req.Request)
	validateItem(x.input, input.RequestItems, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.BatchGetItemExpect = append(e.dynaMock.BatchGetItemExpect[:0], e.dynaMock.BatchGetItemExpect[1:]...)

	req.Data = x.output

	return req
}
