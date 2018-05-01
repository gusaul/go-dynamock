package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"reflect"
)

func (e *DeleteItemExpectation) ToTable(table string) *DeleteItemExpectation {
	e.table = &table
	return e
}

func (e *DeleteItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *DeleteItemExpectation {
	e.key = keys
	return e
}

func (e *DeleteItemExpectation) WillReturns(res dynamodb.DeleteItemOutput) *DeleteItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	if len(e.dynaMock.DeleteItemExpect) > 0 {
		x := e.dynaMock.DeleteItemExpect[0] //get first element of expectation

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
		e.dynaMock.DeleteItemExpect = append(e.dynaMock.DeleteItemExpect[:0], e.dynaMock.DeleteItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Delete Item Expectation Not Found")
}
