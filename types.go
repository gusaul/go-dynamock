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
		GetItemExpect      []GetItemExpectation
		BatchGetItemExpect []BatchGetItemExpectation
		UpdateItemExpect   []UpdateItemExpectation
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

	AnyValue struct{}
)
