package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Name - method for set Name expectation
func (e *CreateTableExpectation) Name(table string) *CreateTableExpectation {
	e.table = &table
	return e
}

// KeySchema - method for set KeySchema expectation
func (e *CreateTableExpectation) KeySchema(keySchema []dynamodb.KeySchemaElement) *CreateTableExpectation {
	e.keySchema = keySchema
	return e
}

// WillReturn - method for set desired result
func (e *CreateTableExpectation) WillReturn(res dynamodb.CreateTableResponse) *CreateTableExpectation {
	e.output = res.CreateTableOutput
	return e
}

// CreateTableRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) CreateTableRequest(input *dynamodb.CreateTableInput) dynamodb.CreateTableRequest {
	req := dynamodb.CreateTableRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.CreateTableExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.CreateTableExpect[0]

	validateInput(input, req.Request)
	validateTable(x.table, input.TableName, req.Request)
	validateItem(x.keySchema, input.KeySchema, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.CreateTableExpect = append(e.dynaMock.CreateTableExpect[:0], e.dynaMock.CreateTableExpect[1:]...)

	req.Data = x.output

	return req
}
