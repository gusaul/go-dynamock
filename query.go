package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table - method for set Table expectation
func (e *QueryExpectation) Table(table string) *QueryExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *QueryExpectation) WillReturns(res dynamodb.QueryOutput) *QueryExpectation {
	e.output = &res
	return e
}

// Query - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	if len(e.dynaMock.QueryExpect) > 0 {
		x := e.dynaMock.QueryExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.QueryOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.QueryExpect = append(e.dynaMock.QueryExpect[:0], e.dynaMock.QueryExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.QueryOutput{}, fmt.Errorf("Query Table Expectation Not Found")
}

// QueryWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, options ...request.Option) (*dynamodb.QueryOutput, error) {
	if len(e.dynaMock.QueryExpect) > 0 {
		x := e.dynaMock.QueryExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.QueryOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.QueryExpect = append(e.dynaMock.QueryExpect[:0], e.dynaMock.QueryExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.QueryOutput{}, fmt.Errorf("Query Table With Context Expectation Not Found")
}
