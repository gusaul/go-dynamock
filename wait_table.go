package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table - method for set Table expectation
func (e *WaitTableExistExpectation) Table(table string) *WaitTableExistExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *WaitTableExistExpectation) WillReturns(err error) *WaitTableExistExpectation {
	e.err = err
	return e
}

// WaitUntilTableExists - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
	if len(e.dynaMock.WaitTableExistExpect) > 0 {
		x := e.dynaMock.WaitTableExistExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.WaitTableExistExpect = append(e.dynaMock.WaitTableExistExpect[:0], e.dynaMock.WaitTableExistExpect[1:]...)

		return x.err
	}

	return fmt.Errorf("Wait Table Exist Expectation Not Found")
}
