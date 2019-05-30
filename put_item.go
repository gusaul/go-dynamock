package dynamock

import (
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *PutItemExpectation) ToTable(table string) *PutItemExpectation {
	e.table = &table
	return e
}

// WithItems - method for set Items expectation
func (e *PutItemExpectation) WithItems(item map[string]dynamodb.AttributeValue) *PutItemExpectation {
	e.item = item
	return e
}

// WillReturn - method for set desired result
func (e *PutItemExpectation) WillReturn(res dynamodb.PutItemResponse) *PutItemExpectation {
	e.output = res.PutItemOutput
	return e
}

func (e *MockDynamoDB) PutItemRequest(input *dynamodb.PutItemInput) dynamodb.PutItemRequest {
	req := dynamodb.PutItemRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.PutItemExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.PutItemExpect[0]

	if x.table != nil {
		if *x.table != *input.TableName {
			req.Error = ErrTableExpectationMismatch

			return req
		}
	}

	if x.item != nil {
		if !reflect.DeepEqual(x.item, input.Item) {
			req.Error = ErrKeyExpectationMismatch

			return req
		}
	}

	e.dynaMock.PutItemExpect = append(e.dynaMock.PutItemExpect[:0], e.dynaMock.PutItemExpect[1:]...)

	req.Data = x.output

	return req
}
