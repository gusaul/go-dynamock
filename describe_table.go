package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table - method for set Table expectation
func (e *DescribeTableExpectation) Table(table string) *DescribeTableExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *DescribeTableExpectation) WillReturns(res dynamodb.DescribeTableOutput) *DescribeTableExpectation {
	e.output = &res
	return e
}

// DescribeTable - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	if len(e.dynaMock.DescribeTableExpect) > 0 {
		x := e.dynaMock.DescribeTableExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.DescribeTableExpect = append(e.dynaMock.DescribeTableExpect[:0], e.dynaMock.DescribeTableExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Describe Table Expectation Not Found")
}
