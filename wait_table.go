package dynamock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Table - method for set Table expectation
func (e *WaitTableExistExpectation) Table(table string) *WaitTableExistExpectation {
	e.table = &table
	return e
}

// WillReturn - method for set desired result
func (e *WaitTableExistExpectation) WillReturn(err error) *WaitTableExistExpectation {
	e.err = err

	return e
}

// WaitUntilTableExists - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) WaitUntilTableExists(ctx context.Context, input *dynamodb.DescribeTableInput, opt ...aws.WaiterOption) error {
	if len(e.dynaMock.WaitTableExistExpect) == 0 {
		return ErrNoExpectation
	}

	x := e.dynaMock.WaitTableExistExpect[0]

	if x.table == nil {
		return ErrNoTable
	}

	if aws.StringValue(x.table) != aws.StringValue(input.TableName) {
		return ErrTableExpectationMismatch
	}

	e.dynaMock.WaitTableExistExpect = append(e.dynaMock.WaitTableExistExpect[:0], e.dynaMock.WaitTableExistExpect[1:]...)

	return nil
}
