package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"reflect"
)

func (e *PutItemExpectation) ToTable(table string) *PutItemExpectation {
	e.table = &table
	return e
}

func (e *PutItemExpectation) WithItems(item map[string]*dynamodb.AttributeValue) *PutItemExpectation {
	e.item = item
	return e
}

func (e *PutItemExpectation) WillReturns(res dynamodb.PutItemOutput) *PutItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if len(e.dynaMock.PutItemExpect) > 0 {
		x := e.dynaMock.PutItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.item != nil {
			if !reflect.DeepEqual(x.item, input.Item) {
				return nil, fmt.Errorf("Expect item %+v but found item %+v", x.item, input.Item)
			}
		}

		// delete first element of expectation
		e.dynaMock.PutItemExpect = append(e.dynaMock.PutItemExpect[:0], e.dynaMock.PutItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Put Item Expectation Not Found")
}
