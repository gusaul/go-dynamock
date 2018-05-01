package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"reflect"
)

func (e *CreateTableExpectation) Name(table string) *CreateTableExpectation {
	e.table = &table
	return e
}

func (e *CreateTableExpectation) KeySchema(keySchema []*dynamodb.KeySchemaElement) *CreateTableExpectation {
	e.keySchema = keySchema
	return e
}

func (e *CreateTableExpectation) WillReturns(res dynamodb.CreateTableOutput) *CreateTableExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	if len(e.dynaMock.CreateTableExpect) > 0 {
		x := e.dynaMock.CreateTableExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.keySchema != nil {
			if !reflect.DeepEqual(x.keySchema, input.KeySchema) {
				return nil, fmt.Errorf("Expect keySchema %+v but found keySchema %+v", x.keySchema, input.KeySchema)
			}
		}

		// delete first element of expectation
		e.dynaMock.CreateTableExpect = append(e.dynaMock.CreateTableExpect[:0], e.dynaMock.CreateTableExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Create Table Expectation Not Found")
}
