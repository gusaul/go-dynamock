package dynamock

import (
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *UpdateItemExpectation) Table(table string) *UpdateItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *UpdateItemExpectation) WithKeys(keys map[string]dynamodb.AttributeValue) *UpdateItemExpectation {
	e.key = keys
	return e
}

// Updates - method for set Updates expectation
func (e *UpdateItemExpectation) Updates(attrs map[string]dynamodb.AttributeValueUpdate) *UpdateItemExpectation {
	e.attributeUpdates = attrs
	return e
}

// WillReturn - method for set desired result
func (e *UpdateItemExpectation) WillReturn(res dynamodb.UpdateItemResponse) *UpdateItemExpectation {
	e.output = res.UpdateItemOutput
	return e
}

// UpdateItemRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItemRequest(input *dynamodb.UpdateItemInput) dynamodb.UpdateItemRequest {
	req := dynamodb.UpdateItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.UpdateItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.UpdateItemExpect[0]

	if x.table != nil {
		if *x.table != *input.TableName {
			req.Error = ErrTableExpectationMismatch

			return req
		}
	}

	if x.attributeUpdates != nil {
		if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
			req.Error = ErrKeyExpectationMismatch

			return req
		}
	}

	e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

	req.Data = x.output

	return req
}
