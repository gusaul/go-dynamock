package dynamock

import (
	"reflect"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// MyDynamo struct hold dynamodb connection
type MyDynamo struct {
	Db dynamodbiface.DynamoDBAPI
}

// Dyna - object from MyDynamo
var Dyna *MyDynamo

func TestOutOfOrderOperations_GetItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		Dyna := new(MyDynamo)
		var mock *DynaMock
		Dyna.Db, mock = New()

		fooOut := dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"name": {
					S: aws.String("foo"),
				},
			},
		}
		mock.ExpectGetItem().ToTable("foo").WillReturns(fooOut)

		barOut :=
			dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"name": {
						S: aws.String("bar"),
					},
				},
			}
		mock.ExpectGetItem().ToTable("bar").WillReturns(barOut)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo, err := Dyna.Db.GetItem(
				&dynamodb.GetItemInput{
					Key: map[string]*dynamodb.AttributeValue{
						"name": {
							N: aws.String("foo"),
						},
					},
					TableName: aws.String("foo"),
				},
			)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(foo.Item["name"].S, fooOut.Item["name"].S) {
				t.Errorf("foo failed expected %v got %v", foo.Item["name"].S, fooOut.Item["name"].S)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			bar, err := Dyna.Db.GetItem(
				&dynamodb.GetItemInput{
					Key: map[string]*dynamodb.AttributeValue{
						"name": {
							N: aws.String("bar"),
						},
					},
					TableName: aws.String("bar"),
				},
			)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(bar.Item["name"].S, barOut.Item["name"].S) {
				t.Errorf("bar failed expected %v got %v", bar.Item["name"].S, barOut.Item["name"].S)
			}
		}()
		wg.Wait()
	}
}

func TestOutOfOrderOperations_GetItem_RaceCondition(t *testing.T) {
	Dyna := new(MyDynamo)
	var mock *DynaMock
	Dyna.Db, mock = New()

	fooOut := dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String("foo"),
			},
		},
	}
	for i := 0; i < 10; i += 1 {
		mock.ExpectGetItem().ToTable("foo").WillReturns(fooOut)
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i += 1 {
		go func() {
			defer wg.Done()
			foo, err := Dyna.Db.GetItem(
				&dynamodb.GetItemInput{
					Key: map[string]*dynamodb.AttributeValue{
						"name": {
							N: aws.String("foo"),
						},
					},
					TableName: aws.String("foo"),
				},
			)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(foo.Item["name"].S, fooOut.Item["name"].S) {
				t.Errorf("foo failed expected %v got %v", foo.Item["name"].S, fooOut.Item["name"].S)
			}
		}()
	}

	wg.Wait()
	if len(mock.GetItemExpect) != 0 {
		t.Fatal("We should have exausted all the expectations")
	}
}
