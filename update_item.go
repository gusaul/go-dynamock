package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *UpdateItemExpectation) ToTable(table string) *UpdateItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *UpdateItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *UpdateItemExpectation {
	e.key = keys
	return e
}

// WithConditionExpression - method for setting a ConditionExpression expectation
func (e *UpdateItemExpectation) WithConditionExpression(expr *string) *UpdateItemExpectation {
	e.conditionExpression = expr
	return e
}

// WithExpressionAttributeNames - method for setting a ExpressionAttributeNames expectation
func (e *UpdateItemExpectation) WithExpressionAttributeNames(names map[string]*string) *UpdateItemExpectation {
	e.expressionAttributeNames = names
	return e
}

// WithExpressionAttributeValues - method for setting a ExpressionAttributeValues expectation
func (e *UpdateItemExpectation) WithExpressionAttributeValues(attrs map[string]*dynamodb.AttributeValue) *UpdateItemExpectation {
	e.expressionAttributeValues = attrs
	return e
}

// WithUpdateExpression - method for setting a UpdateExpression expectation
func (e *UpdateItemExpectation) WithUpdateExpression(expr *string) *UpdateItemExpectation {
	e.updateExpression = expr
	return e
}

// Updates - method for set Updates expectation
func (e *UpdateItemExpectation) Updates(attrs map[string]*dynamodb.AttributeValueUpdate) *UpdateItemExpectation {
	e.attributeUpdates = attrs
	return e
}

// WillReturns - method for set desired result
func (e *UpdateItemExpectation) WillReturns(res dynamodb.UpdateItemOutput) *UpdateItemExpectation {
	e.output = &res
	return e
}

// UpdateItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		if x.conditionExpression != nil {
			if !reflect.DeepEqual(x.conditionExpression, input.ConditionExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.conditionExpression, input.ConditionExpression)
			}
		}

		if x.expressionAttributeNames != nil {
			if !reflect.DeepEqual(x.expressionAttributeNames, input.ExpressionAttributeNames) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.expressionAttributeNames, input.ExpressionAttributeNames)
			}
		}

		if x.expressionAttributeValues != nil {
			if !reflect.DeepEqual(x.expressionAttributeValues, input.ExpressionAttributeValues) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.expressionAttributeValues, input.ExpressionAttributeValues)
			}
		}

		if x.updateExpression != nil {
			if !reflect.DeepEqual(x.updateExpression, input.UpdateExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.updateExpression, input.UpdateExpression)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item Expectation Not Found")
}

// UpdateItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opt ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		if x.conditionExpression != nil {
			if !reflect.DeepEqual(x.conditionExpression, input.ConditionExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.conditionExpression, input.ConditionExpression)
			}
		}

		if x.expressionAttributeNames != nil {
			if !reflect.DeepEqual(x.expressionAttributeNames, input.ExpressionAttributeNames) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.expressionAttributeNames, input.ExpressionAttributeNames)
			}
		}

		if x.expressionAttributeValues != nil {
			if !reflect.DeepEqual(x.expressionAttributeValues, input.ExpressionAttributeValues) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.expressionAttributeValues, input.ExpressionAttributeValues)
			}
		}

		if x.updateExpression != nil {
			if !reflect.DeepEqual(x.updateExpression, input.UpdateExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.updateExpression, input.UpdateExpression)
			}
		}
		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item With Context Expectation Not Found")
}
