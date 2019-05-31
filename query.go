package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *QueryExpectation) Table(table string) *QueryExpectation {
	e.table = &table
	return e
}

// WillReturn - method for set desired result
func (e *QueryExpectation) WillReturn(res dynamodb.QueryResponse) *QueryExpectation {
	e.output = res.QueryOutput
	return e
}

// QueryRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) QueryRequest(input *dynamodb.QueryInput) dynamodb.QueryRequest {
	req := dynamodb.QueryRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.QueryExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.QueryExpect[0]

	validateInput(input, req.Request)
	validateTable(x.table, input.TableName, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.QueryExpect = append(e.dynaMock.QueryExpect[:0], e.dynaMock.QueryExpect[1:]...)

	req.Data = x.output

	return req
}
