[![GoDoc](https://godoc.org/github.com/gusaul/go-dynamock?status.png)](https://godoc.org/github.com/gusaul/go-dynamock)
# go-dynamock
Amazon Dynamo DB Mock Driver for Golang to Test Database Interactions

## Install
```
go get github.com/gusaul/go-dynamock
```

## Examples Usage
Visit [godoc](https://godoc.org/github.com/gusaul/go-dynamock) for general examples and public api reference.

### DynamoDB configuration
First of all, change the dynamodb configuration to use the ***dynamodb interface***. see code below:
``` go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MyDynamo struct {
    Db dynamodbiface.DynamoDBAPI
}

var Dyna *MyDynamo

func ConfigureDynamoDB() {
    Dyna = new(MyDynamo)
    awsSession, _ := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-2")})
    var svc *dynamodb.DynamoDB = dynamodb.New(awsSession)
    Dyna.Db = dynamodbiface.DynamoDBAPI(svc)
}
```
the purpose of code above is to make your dynamoDB object can be mocked by ***dynamock*** through the dynamodbiface.

### Something you may wanna test
``` go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetName(id string) (*string, error) {
    parameter := &dynamodb.GetItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "id" : {
                N: aws.String(id)
            },
        },
        TableName: "employee"
    }

    response, err := Dyna.Db.GetItem(parameter)
    if err != nil {
        return nil, err
    }

    name := response["name"].S
    return name, nil
}
```

### Test with DynaMock
``` go
package main

import (
    "testing"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/gusaul/go-dynamock"
)

var mock *dynamock.DynaMock

func init() {
    Dyna.db, mock = dynamock.New()
}

func TestGetName(t *testing.T) {
    expectKey := map[string]*dynamodb.AttributeValue{
        "id" : {
            N: aws.String("1")
        },
    }

    expectedResult := aws.String("jaka")
    result := dynamodb.GetItemOutput{
        Item: map[string]*dynamodb.AttributeValue{
            "name": {
                S: expectedResult,
            },
        },
    }

    //lets start dynamock in action
    mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturns(result)

    actualResult, _ := GetName("1")
    if actualResult != expectedResult {
        t.Errorf("Test Fail")
    }
}
```
if you just wanna expect the table
``` go
mock.ExpectGetItem().ToTable("employee").WillReturns(result)
```
or maybe you didn't care with any arguments, you just need to determine the result
``` go
mock.ExpectGetItem().WillReturns(result)
```
and you can do multiple expectations at once, then the expectation will be executed sequentially.
``` go
mock.ExpectGetItem().WillReturns(resultOne)
mock.ExpectUpdateItem().WillReturns(resultTwo)
mock.ExpectGetItem().WillReturns(resultThree)

/* Result
the first call of GetItem will return resultOne
the second call of GetItem will return resultThree
and the only call of UpdateItem will return resultTwo */
```
### Currently Supported Functions
``` go
CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
GetItemWithContext(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error)
PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
UpdateItemWithContext(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error)
DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
BatchGetItemWithContext(aws.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error)
BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
WaitUntilTableExists(*dynamodb.DescribeTableInput) error
Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
```
## Contributions

Feel free to open a pull request. Note, if you wish to contribute an extension to public (exported methods or types) -
please open an issue before, to discuss whether these changes can be accepted. All backward incompatible changes are
and will be treated cautiously

## License

The [three clause BSD license](http://en.wikipedia.org/wiki/BSD_licenses)
