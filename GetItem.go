package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (e *GetItemExpectation) ToTable(table string) *GetItemExpectation {
	e.table = &table
	return e
}

func (e *GetItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *GetItemExpectation {
	e.key = keys
	return e
}

func (e *GetItemExpectation) ExpectReturns(res dynamodb.GetItemOutput) *GetItemExpectation {
	e.output = &res
	return e
}
