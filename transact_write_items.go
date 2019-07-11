package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table - method for set Table expectation
func (e *TransactWriteItemsExpectation) Table(table string) *TransactWriteItemsExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *TransactWriteItemsExpectation) WillReturns(res dynamodb.TransactWriteItemsOutput) *TransactWriteItemsExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error){
	if len(e.dynaMock.TransactWriteItemsExpect) > 0 {
		x := e.dynaMock.TransactWriteItemsExpect[0] //get first element of expectation

		foundTable := false
		if x.table != nil {
			for _, item := range input.TransactItems {
				if x.table == item.Update.TableName || x.table == item.Put.TableName ||
					x.table == item.Delete.TableName {
					foundTable = true
				}
			}

			if foundTable == false {
				return nil, fmt.Errorf("Expect table %s not found", *x.table)
			}
		}

		// delete first element of expectation
		e.dynaMock.TransactWriteItemsExpect = append(e.dynaMock.TransactWriteItemsExpect[:0],
			e.dynaMock.TransactWriteItemsExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Transact Write Items Table Expectation Not Found")
}