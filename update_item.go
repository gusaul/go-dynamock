package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (e *UpdateItemExpectation) ToTable(table string) *UpdateItemExpectation {
	e.table = &table
	return e
}

func (e *UpdateItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *UpdateItemExpectation {
	e.key = keys
	return e
}

func (e *UpdateItemExpectation) Updates(attrs map[string]*dynamodb.AttributeValueUpdate) *UpdateItemExpectation {
	e.attributeUpdates = attrs
	return e
}

func (e *UpdateItemExpectation) WillReturns(res dynamodb.UpdateItemOutput) *UpdateItemExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

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

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return nil, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Update Item Expectation Not Found")
}

func (e *MockDynamoDB) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

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

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return nil, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Update Item With Context Expectation Not Found")
}
