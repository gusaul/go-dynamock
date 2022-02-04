package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *GetItemExpectation) ToTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *GetItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

// WillReturns - method for set desired result
func (e *GetItemExpectation) WillReturns(res dynamodb.GetItemOutput) *GetItemExpectation {
	e.output = &res
	return e
}

// GetItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for i, x := range e.dynaMock.GetItemExpect {
		// If we the expectation doesn't contain a table or it matches the input match
		if x.table == nil || reflect.DeepEqual(x.table, input.TableName) {
			// If we got no keys on the expectation match on any input
			if len(x.key) == 0 || reflect.DeepEqual(x.key, input.Key) {
				end := i + 1
				if len(e.dynaMock.GetItemExpect) == i {
					end = i
				}
				e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:i], e.dynaMock.GetItemExpect[end:]...)
				return x.output, nil
			}
		}
	}

	return &dynamodb.GetItemOutput{}, fmt.Errorf("Expected input with table name %+v and key %+v. Could not find within expectations", input.TableName, input.Key)
}

// GetItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opt ...request.Option) (*dynamodb.GetItemOutput, error) {
	if len(e.dynaMock.GetItemExpect) > 0 {
		x := e.dynaMock.GetItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.GetItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.GetItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:0], e.dynaMock.GetItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.GetItemOutput{}, fmt.Errorf("Get Item With Context Expectation Not Found")
}
