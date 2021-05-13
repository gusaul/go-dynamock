package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"regexp"
	"sort"
	"strings"
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

func (e *UpdateItemExpectation) WithEquivalentUpdateExpression(expr *string) *UpdateItemExpectation {
	parsed := parseUpdateExpression(*expr)
	e.equivalentUpdateExpression = &parsed
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
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect AttributeUpdates: %+v but got: %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		if x.conditionExpression != nil {
			if !reflect.DeepEqual(x.conditionExpression, input.ConditionExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect ConditionExpressions: %+v but got: %+v", x.conditionExpression, input.ConditionExpression)
			}
		}

		if x.expressionAttributeNames != nil {
			if !reflect.DeepEqual(x.expressionAttributeNames, input.ExpressionAttributeNames) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect ExpressionAttributeNames: %+v but got: %+v", x.expressionAttributeNames, input.ExpressionAttributeNames)
			}
		}

		if x.expressionAttributeValues != nil {
			if !reflect.DeepEqual(x.expressionAttributeValues, input.ExpressionAttributeValues) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect ExpressionAttributeValues: %+v but got: %+v", x.expressionAttributeValues, input.ExpressionAttributeValues)
			}
		}

		if x.updateExpression != nil {
			if !reflect.DeepEqual(x.updateExpression, input.UpdateExpression) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("expect UpdateExpression: %+v but got: %+v", x.updateExpression, input.UpdateExpression)
			}
		}

		if x.equivalentUpdateExpression != nil {
			inputExpr := parseUpdateExpression(*input.UpdateExpression)
			err := x.equivalentUpdateExpression.CheckIsEquivalentTo(&inputExpr)
			if err != nil {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("non-equivalent update expressions found: %v", err)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item Expectation Not Found")
}

type parsedUpdateExpression struct {
	ADDExpressions    []pathValueExpression
	DELETEExpressions []pathValueExpression
	REMOVEExpressions []pathExpression
	SETExpressions    []pathValueExpression
}

func (p *parsedUpdateExpression) CheckIsEquivalentTo(other *parsedUpdateExpression) error {
	sort.Slice(p.ADDExpressions, func(i, j int) bool {
		return p.ADDExpressions[i].path < p.ADDExpressions[j].path
	})
	sort.Slice(other.ADDExpressions, func(i, j int) bool {
		return other.ADDExpressions[i].path < other.ADDExpressions[j].path
	})
	if !reflect.DeepEqual(p.ADDExpressions, other.ADDExpressions) {
		return fmt.Errorf("ADDExpressions do not match, %v != %v", p.ADDExpressions, other.ADDExpressions)
	}
	sort.Slice(p.DELETEExpressions, func(i, j int) bool {
		return p.DELETEExpressions[i].path < p.DELETEExpressions[j].path
	})
	sort.Slice(other.DELETEExpressions, func(i, j int) bool {
		return other.DELETEExpressions[i].path < other.DELETEExpressions[j].path
	})
	if !reflect.DeepEqual(p.DELETEExpressions, other.DELETEExpressions) {
		return fmt.Errorf("DELETEExpressions do not match, %v != %v", p.DELETEExpressions, other.DELETEExpressions)
	}
	sort.Slice(p.REMOVEExpressions, func(i, j int) bool {
		return p.REMOVEExpressions[i].path < p.REMOVEExpressions[j].path
	})
	sort.Slice(other.REMOVEExpressions, func(i, j int) bool {
		return other.REMOVEExpressions[i].path < other.REMOVEExpressions[j].path
	})
	if !reflect.DeepEqual(p.REMOVEExpressions, other.REMOVEExpressions) {
		return fmt.Errorf("REMOVEExpressions do not match, %v != %v", p.REMOVEExpressions, other.REMOVEExpressions)
	}
	sort.Slice(p.SETExpressions, func(i, j int) bool {
		return p.SETExpressions[i].path < p.SETExpressions[j].path
	})
	sort.Slice(other.SETExpressions, func(i, j int) bool {
		return other.SETExpressions[i].path < other.SETExpressions[j].path
	})
	if !reflect.DeepEqual(p.SETExpressions, other.SETExpressions) {
		return fmt.Errorf("SETExpressions do not match, %v != %v", p.SETExpressions, other.SETExpressions)
	}
	return nil
}

type operation string

const (
	ADD    operation = "ADD"
	DELETE operation = "DELETE"
	REMOVE operation = "REMOVE"
	SET    operation = "SET"
)

type operationIndexTuple struct {
	Index     int
	Operation operation
}

type pathValueExpression struct {
	path  string
	value string
}

type pathExpression struct {
	path string
}

func mustExtractPathValueExpressions(operation operation, expr string) []pathValueExpression {
	var re *regexp.Regexp
	var subMatchRe *regexp.Regexp
	var result []pathValueExpression
	switch operation {
	case ADD:
		re = regexp.MustCompile(`ADD\s+((\S+\s+[\w:#]+\s*,?\s*)+)`)
		subMatchRe = regexp.MustCompile(`(\S+)\s+([\w:#]+)\s*,?\s*`)
	case DELETE:
		re = regexp.MustCompile(`DELETE\s+((\S+\s+[\w:#]+\s*,?\s*)+)`)
		subMatchRe = regexp.MustCompile(`(\S+)\s+([\w:#]+)\s*,?\s*`)
	case SET:
		// SET operations allow the value to two operands with a + or - in between them
		// SET operations allow the operand to be a function such as `SET #ri = list_append(#ri, :vals)`
		re = regexp.MustCompile(`SET\s+((\S+\s*=\s*[\w:#\(\)\+-,\s=]+\s*,?\s*)+)`)
		subMatchRe = regexp.MustCompile(`(\S+)\s*=\s*([\w:#\+-]+(\([\w\s,:#]*\))?)\s*,?\s*`)
	}
	if re == nil {
		return result
	}
	if subMatchRe == nil {
		return result
	}

	subMatches := re.FindStringSubmatch(expr)

	if subMatches == nil {
		return result
	}

	pairMatches := subMatchRe.FindAllStringSubmatch(subMatches[1], -1)
	if pairMatches == nil {
		return result
	}
	for _, subMatch := range pairMatches {
		result = append(result, pathValueExpression{subMatch[1], subMatch[2]})
	}
	return result
}

func extractAddPathValuePairs(addExpr string) []pathValueExpression {
	return mustExtractPathValueExpressions(ADD, addExpr)
}

func extractDeletePathValuePairs(deleteExpr string) []pathValueExpression {
	return mustExtractPathValueExpressions(DELETE, deleteExpr)
}

func extractRemovePath(removeExpr string) []pathExpression {
	re := regexp.MustCompile(`REMOVE\s+(([\w:#\[\]]+\s*,?\s*)+)`)
	subMatchRe := regexp.MustCompile(`\s*([\w:#\[\]]+)\s*,?\s*`)
	subMatches := re.FindStringSubmatch(removeExpr)
	var result []pathExpression
	if subMatches == nil {
		return result
	}
	pairMatches := subMatchRe.FindAllStringSubmatch(subMatches[1], -1)
	if pairMatches == nil {
		return result
	}
	for _, subMatch := range pairMatches {
		result = append(result, pathExpression{subMatch[1]})
	}
	return result
}

func extractSetPathValuePairs(setExpr string) []pathValueExpression {
	return mustExtractPathValueExpressions(SET, setExpr)
}

func parseUpdateExpression(updateExpression string) parsedUpdateExpression {
	addOp := operationIndexTuple{strings.Index(updateExpression, "ADD"), ADD}
	deleteOp := operationIndexTuple{strings.Index(updateExpression, "DELETE"), DELETE}
	removeOp := operationIndexTuple{strings.Index(updateExpression, "REMOVE"), REMOVE}
	setOp := operationIndexTuple{strings.Index(updateExpression, "SET"), SET}

	ops := []operationIndexTuple{
		addOp,
		deleteOp,
		removeOp,
		setOp,
	}
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].Index < ops[j].Index
	})

	result := parsedUpdateExpression{}
	for opIdx, op := range ops {
		if op.Index < 0 {
			// op.Index should be -1 for operations that are not present in an update expression
			continue
		}
		// get the substring for the operation
		var substr string
		if opIdx+1 < len(ops) {
			// We don't need to worry about the case where (opIdx+1).Index is -1, because we're iterating through a
			// ascending sorted array.
			substr = updateExpression[op.Index:ops[opIdx+1].Index]
		} else {
			substr = updateExpression[op.Index:]
		}
		// apply the operation specific parsing
		switch op.Operation {
		case ADD:
			result.ADDExpressions = extractAddPathValuePairs(substr)
		case DELETE:
			result.DELETEExpressions = extractDeletePathValuePairs(substr)
		case REMOVE:
			result.REMOVEExpressions = extractRemovePath(substr)
		case SET:
			result.SETExpressions = extractSetPathValuePairs(substr)
		}
	}
	return result
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
		if x.equivalentUpdateExpression != nil {
			inputExpr := parseUpdateExpression(*input.UpdateExpression)
			err := x.equivalentUpdateExpression.CheckIsEquivalentTo(&inputExpr)
			if err != nil {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("non-equivalent update expressions found: %v", err)
			}
		}
		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item With Context Expectation Not Found")
}
