package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *GetItemExpectation) Table(table string) *GetItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *GetItemExpectation) WithKeys(keys map[string]dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

// WillReturn - method for set desired result
func (e *GetItemExpectation) WillReturn(res dynamodb.GetItemResponse) *GetItemExpectation {
	e.output = res.GetItemOutput
	return e
}

// GetItemRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) GetItemRequest(input *dynamodb.GetItemInput) dynamodb.GetItemRequest {
	req := dynamodb.GetItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.GetItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.GetItemExpect[0]

	validateInput(input, req.Request)
	validateTable(input.TableName, x.table, req.Request)
	validateKey(input.Key, x.key, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:0], e.dynaMock.GetItemExpect[1:]...)

	req.Data = x.output

	return req
}
