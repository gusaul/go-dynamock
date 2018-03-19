package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (e *QueryExpectation) Table(table string) *QueryExpectation {
	e.table = &table
	return e
}

func (e *QueryExpectation) WillReturns(res dynamodb.QueryOutput) *QueryExpectation {
	e.output = &res
	return e
}

func (e *MockDynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	if len(e.dynaMock.QueryExpect) > 0 {
		x := e.dynaMock.QueryExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return nil, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		// delete first element of expectation
		e.dynaMock.QueryExpect = append(e.dynaMock.QueryExpect[:0], e.dynaMock.QueryExpect[1:]...)

		return x.output, nil
	}

	return nil, fmt.Errorf("Query Table Expectation Not Found")
}
