package dynamock

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *ScanExpectation) Table(table string) *ScanExpectation {
	e.table = &table
	return e
}

// WillReturn - method for set desired result
func (e *ScanExpectation) WillReturn(res dynamodb.ScanResponse) *ScanExpectation {
	e.output = res.ScanOutput
	return e
}

// ScanRequest - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) ScanRequest(input *dynamodb.ScanInput) dynamodb.ScanRequest {
	req := dynamodb.ScanRequest{
		Request: &aws.Request{
			HTTPRequest: &http.Request{},
		},
	}

	if len(e.dynaMock.ScanExpect) == 0 {
		req.Error = ErrNoExpectation

		return req
	}

	x := e.dynaMock.ScanExpect[0]

	validateInput(input, req.Request)
	validateTable(x.table, input.TableName, req.Request)
	if req.Error != nil {
		return req
	}

	e.dynaMock.ScanExpect = append(e.dynaMock.ScanExpect[:0], e.dynaMock.ScanExpect[1:]...)

	req.Data = x.output

	return req
}
