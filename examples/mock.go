package examples

import (
	dynamock "github.com/groovili/go-dynamock"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
)

// FakeDynamo struct hold dynamodb connection
type FakeDynamo struct {
	DB dynamodbiface.ClientAPI
}

// Fake - object from MyDynamo
var Fake *FakeDynamo

var Mock *dynamock.DynaMock
