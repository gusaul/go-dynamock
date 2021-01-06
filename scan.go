package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Table - method for set Table expectation
func (e *ScanExpectation) Table(table string) *ScanExpectation {
	e.table = &table
	return e
}

// WillReturns - method for set desired result
func (e *ScanExpectation) WillReturns(res dynamodb.ScanOutput) *ScanExpectation {
	e.output = &res
	return e
}

// Scan - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if len(e.dynaMock.ScanExpect) > 0 {
		x := e.dynaMock.ScanExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.ScanOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.ScanExpect = append(e.dynaMock.ScanExpect[:0], e.dynaMock.ScanExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.ScanOutput{}, fmt.Errorf("Scan Table Expectation Not Found")
}

// ScanWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	if len(e.dynaMock.ScanExpect) > 0 {
		x := e.dynaMock.ScanExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.ScanOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.ScanExpect = append(e.dynaMock.ScanExpect[:0], e.dynaMock.ScanExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.ScanOutput{}, fmt.Errorf("Scan Table With Context Expectation Not Found")
}

// ScanPages - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) ScanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	if len(e.dynaMock.ScanExpect) > 0 {
		x := e.dynaMock.ScanExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.ScanExpect = append(e.dynaMock.ScanExpect[:0], e.dynaMock.ScanExpect[1:]...)

		fn(x.output, true)
		return nil
	}

	return fmt.Errorf("Scan Table By Page Expectation Not Found")
}

// ScanPagesWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	if len(e.dynaMock.ScanExpect) > 0 {
		x := e.dynaMock.ScanExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.ScanExpect = append(e.dynaMock.ScanExpect[:0], e.dynaMock.ScanExpect[1:]...)

		fn(x.output, true)
		return nil
	}

	return fmt.Errorf("Scan Table By Page With Context Expectation Not Found")
}
