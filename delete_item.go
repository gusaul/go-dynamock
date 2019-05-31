package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *DeleteItemExpectation) Table(table string) *DeleteItemExpectation {
	e.table = &table

	return e
}

// WillReturn - method for set desired result
func (e *DeleteItemExpectation) WillReturn(res dynamodb.DeleteItemResponse) *DeleteItemExpectation {
	e.output = res.DeleteItemOutput

	return e
}

// WithKeys - method for set Keys expectation
func (e *DeleteItemExpectation) WithKeys(keys map[string]dynamodb.AttributeValue) *DeleteItemExpectation {
	e.key = keys

	return e
}

// DeleteItemRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) DeleteItemRequest(input *dynamodb.DeleteItemInput) dynamodb.DeleteItemRequest {
	req := dynamodb.DeleteItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.DeleteItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.DeleteItemExpect[0]

	validateInput(input, req.Request)
	validateTable(input.TableName, x.table, req.Request)
	validateKey(input.Key, x.key, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.DeleteItemExpect = append(e.dynaMock.DeleteItemExpect[:0], e.dynaMock.DeleteItemExpect[1:]...)

	req.Data = x.output

	return req
}
