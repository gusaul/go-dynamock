package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type (
	MockDynamoDB struct {
		dynamodbiface.DynamoDBAPI
		dynaMock *DynaMock
	}

	DynaMock struct {
		GetItemExpect        []GetItemExpectation
		BatchGetItemExpect   []BatchGetItemExpectation
		UpdateItemExpect     []UpdateItemExpectation
		DeleteItemExpect     []DeleteItemExpectation
		BatchWriteItemExpect []BatchWriteItemExpectation
		CreateTableExpect    []CreateTableExpectation
		DescribeTableExpect  []DescribeTableExpectation
	}

	GetItemExpectation struct {
		table  *string
		key    map[string]*dynamodb.AttributeValue
		output *dynamodb.GetItemOutput
	}

	BatchGetItemExpectation struct {
		input  map[string]*dynamodb.KeysAndAttributes
		output *dynamodb.BatchGetItemOutput
	}

	UpdateItemExpectation struct {
		attributeUpdates map[string]*dynamodb.AttributeValueUpdate
		key              map[string]*dynamodb.AttributeValue
		table            *string
		output           *dynamodb.UpdateItemOutput
	}

	DeleteItemExpectation struct {
		key    map[string]*dynamodb.AttributeValue
		table  *string
		output *dynamodb.DeleteItemOutput
	}

	BatchWriteItemExpectation struct {
		input  map[string][]*dynamodb.WriteRequest
		output *dynamodb.BatchWriteItemOutput
	}

	CreateTableExpectation struct {
		keySchema []*dynamodb.KeySchemaElement
		table     *string
		output    *dynamodb.CreateTableOutput
	}

	DescribeTableExpectation struct {
		table  *string
		output *dynamodb.DescribeTableOutput
	}

	AnyValue struct{}
)
