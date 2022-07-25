package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *DeleteItemExpectation) ToTable(table string) *DeleteItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *DeleteItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *DeleteItemExpectation {
	e.key = keys
	return e
}

// WillReturns - method for set desired result
func (e *DeleteItemExpectation) WillReturns(res dynamodb.DeleteItemOutput) *DeleteItemExpectation {
	e.output = &res
	return e
}

// DeleteItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	if len(e.dynaMock.DeleteItemExpect) > 0 {
		x := e.dynaMock.DeleteItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.dynaMock.DeleteItemExpect = append(e.dynaMock.DeleteItemExpect[:0], e.dynaMock.DeleteItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Delete Item Expectation Not Found")
}

// DeleteItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	if len(e.dynaMock.DeleteItemExpect) > 0 {
		x := e.dynaMock.DeleteItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.dynaMock.DeleteItemExpect = append(e.dynaMock.DeleteItemExpect[:0], e.dynaMock.DeleteItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.DeleteItemOutput{}, fmt.Errorf("Delete Item With Context Expectation Not Found")
}
