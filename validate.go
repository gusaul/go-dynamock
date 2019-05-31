package dynamock

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func validateInput(v aws.Validator, r *aws.Request) {
	if err := v.Validate(); err != nil {
		r.Error = ErrInvalidInput
	}
}

func validateKey(wantKey interface{}, gotKey interface{}, r *aws.Request) {
	if wantKey == nil {
		r.Error = ErrNoKey

		return
	}

	if !reflect.DeepEqual(wantKey, gotKey) {
		r.Error = ErrKeyExpectationMismatch

		return
	}
}

func validateItem(wantItem interface{}, gotItem interface{}, r *aws.Request) {
	if wantItem == nil {
		r.Error = ErrNoItem

		return
	}

	if !reflect.DeepEqual(wantItem, gotItem) {
		r.Error = ErrItemExpectationMismatch

		return
	}
}

func validateTable(wantTable *string, gotTable *string, r *aws.Request) {
	if gotTable == nil {
		r.Error = ErrNoTable

		return
	}

	if !(aws.StringValue(gotTable) == aws.StringValue(wantTable)) {
		r.Error = ErrTableExpectationMismatch

		return
	}
}
