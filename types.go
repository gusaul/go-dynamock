package dynamock

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
)

type (
	// MockDynamoDB struct hold DynamoDBAPI implementation and mock object
	MockDynamoDB struct {
		dynamodbiface.ClientAPI
		dynaMock *DynaMock
	}

	// DynaMock mock struct hold all expectation types
	DynaMock struct {
		GetItemExpect        []*GetItemExpectation
		BatchGetItemExpect   []*BatchGetItemExpectation
		UpdateItemExpect     []*UpdateItemExpectation
		PutItemExpect        []*PutItemExpectation
		DeleteItemExpect     []*DeleteItemExpectation
		BatchWriteItemExpect []*BatchWriteItemExpectation
		CreateTableExpect    []*CreateTableExpectation
		DescribeTableExpect  []*DescribeTableExpectation
		WaitTableExistExpect []*WaitTableExistExpectation
		ScanExpect           []*ScanExpectation
		QueryExpect          []*QueryExpectation
	}

	// GetItemExpectation struct hold expectation field, err, and result
	GetItemExpectation struct {
		table  *string
		key    map[string]dynamodb.AttributeValue
		output *dynamodb.GetItemOutput
	}

	// PutItemExpectation struct hold expectation field, err, and result
	PutItemExpectation struct {
		item   map[string]dynamodb.AttributeValue
		table  *string
		output *dynamodb.PutItemOutput
	}

	// BatchGetItemExpectation struct hold expectation field, err, and result
	BatchGetItemExpectation struct {
		input  map[string]dynamodb.KeysAndAttributes
		output *dynamodb.BatchGetItemOutput
	}

	// UpdateItemExpectation struct hold expectation field, err, and result
	UpdateItemExpectation struct {
		attributeUpdates map[string]dynamodb.AttributeValueUpdate
		key              map[string]dynamodb.AttributeValue
		table            *string
		output           *dynamodb.UpdateItemOutput
	}

	// DeleteItemExpectation struct hold expectation field, err, and result
	DeleteItemExpectation struct {
		key    map[string]dynamodb.AttributeValue
		table  *string
		output *dynamodb.DeleteItemOutput
	}

	// BatchWriteItemExpectation struct hold expectation field, err, and result
	BatchWriteItemExpectation struct {
		input  map[string][]dynamodb.WriteRequest
		output *dynamodb.BatchWriteItemOutput
	}

	// CreateTableExpectation struct hold expectation field, err, and result
	CreateTableExpectation struct {
		keySchema []dynamodb.KeySchemaElement
		table     *string
		output    *dynamodb.CreateTableOutput
	}

	// DescribeTableExpectation struct hold expectation field, err, and result
	DescribeTableExpectation struct {
		table  *string
		output *dynamodb.DescribeTableOutput
	}

	// WaitTableExistExpectation struct hold expectation field, err, and result
	WaitTableExistExpectation struct {
		table *string
		err   error
	}

	// ScanExpectation struct hold expectation field, err, and result
	ScanExpectation struct {
		table  *string
		output *dynamodb.ScanOutput
	}

	// QueryExpectation struct hold expectation field, err, and result
	QueryExpectation struct {
		table  *string
		output *dynamodb.QueryOutput
	}
)
