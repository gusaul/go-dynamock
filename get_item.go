package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (e *GetItemExpectation) ToTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

func (e *GetItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

func (e *GetItemExpectation) WillReturns(res dynamodb.GetItemOutput) *GetItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	if len(e.dynaMock.GetItemExpect) > 0 {
		x := e.dynaMock.GetItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return nil, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:0], e.dynaMock.GetItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Get Item Expectation Not Found")
}

func (e *MockDynamoDB) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	if len(e.dynaMock.GetItemExpect) > 0 {
		x := e.dynaMock.GetItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return nil, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		// delete first element of expectation
		e.dynaMock.GetItemExpect = append(e.dynaMock.GetItemExpect[:0], e.dynaMock.GetItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Get Item With Context Expectation Not Found")
}
