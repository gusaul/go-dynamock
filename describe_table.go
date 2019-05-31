package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *DescribeTableExpectation) Table(table string) *DescribeTableExpectation {
	e.table = &table
	return e
}

// WillReturn - method for set desired result
func (e *DescribeTableExpectation) WillReturn(res dynamodb.DescribeTableResponse) *DescribeTableExpectation {
	e.output = res.DescribeTableOutput
	return e
}

// DescribeTableRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) DescribeTableRequest(input *dynamodb.DescribeTableInput) dynamodb.DescribeTableRequest {
	req := dynamodb.DescribeTableRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.DescribeTableExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.DescribeTableExpect[0]

	validateInput(input, req.Request)
	validateTable(x.table, input.TableName, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.DescribeTableExpect = append(e.dynaMock.DescribeTableExpect[:0], e.dynaMock.DescribeTableExpect[1:]...)

	req.Data = x.output

	return req
}
