package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
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

// TransactWriteItems - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error){
	if len(e.dynaMock.TransactWriteItemsExpect) > 0 {
		x := e.dynaMock.TransactWriteItemsExpect[0] //get first element of expectation

		foundTable := false

		if x.table != nil {
			for _, item := range input.TransactItems {
				if (item.Update != nil && x.table == item.Update.TableName) ||
					(item.Put != nil && x.table == item.Put.TableName) ||
					(item.Delete != nil && x.table == item.Delete.TableName) {
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

// TransactWriteItemsWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) TransactWriteItemsWithContext(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, opt ...request.Option) (*dynamodb.TransactWriteItemsOutput, error){
	if len(e.dynaMock.TransactWriteItemsExpect) > 0 {
		x := e.dynaMock.TransactWriteItemsExpect[0] //get first element of expectation

		foundTable := false

		if x.table != nil {
			for _, item := range input.TransactItems {
				if (item.Update != nil && x.table == item.Update.TableName) ||
					(item.Put != nil && x.table == item.Put.TableName) ||
					(item.Delete != nil && x.table == item.Delete.TableName) {
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