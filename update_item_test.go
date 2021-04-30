package dynamock

import (
	"reflect"
	"testing"
)

func Test_extractAddPathValuePairs(t *testing.T) {
	type args struct {
		addExpr string
	}
	tests := []struct {
		name string
		args args
		want []addExpression
	}{
		{
			"no ADD in expression produces empty result",
			args{"SET foo = 3"},
			[]addExpression{},
		},
		{
			"ADD with single pair captures correct pair",
			args{"ADD foobar 3"},
			[]addExpression{{"foobar", "3"}},
		},
		{
			"ADD with single pair and trailing comma captures correct pair",
			args{"ADD foobar 3 ,"},
			[]addExpression{{"foobar", "3"}},
		},
		{
			"ADD with single pair and trailing whitespace captures correct pair",
			args{"ADD foobar 3   "},
			[]addExpression{{"foobar", "3"}},
		},
		{
			"ADD with multiple pairs captures the pairs",
			args{"ADD foobar 3, bazdog 7, chicken 8"},
			[]addExpression{
				{"foobar", "3"},
				{"bazdog", "7"},
				{"chicken", "8"},
			},
		},
		{
			"ADD with multiple pairs and curious whitespace",
			args{"ADD 	  foobar   	3  ,   bazdog 7   ,chicken     8"},
			[]addExpression{
				{"foobar", "3"},
				{"bazdog", "7"},
				{"chicken", "8"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractAddPathValuePairs(tt.args.addExpr)
			if len(got) != len(tt.want) {
				t.Errorf("extractAddPathValuePairs() len = %v, wanted len %v", len(got), len(tt.want))
				return
			}
			for idx, pair := range got {
				if !reflect.DeepEqual(pair, tt.want[idx]) {
					t.Errorf("extractAddPathValuePairs() %vth item got: %v, want: %v", idx, pair, tt.want[idx])
					return
				}
			}
		})
	}
}

func Test_parseUpdateExpression(t *testing.T) {
	type args struct {
		updateExpression string
	}
	tests := []struct {
		name string
		args args
		want parsedUpdateExpression
	}{
		{
			"ADD expression only has only Add expressions",
			args{"ADD foobar 3, dog 76"},
			parsedUpdateExpression{
				ADDExpressions: []addExpression{
					{"foobar", "3"},
					{"dog", "76"},
				},
				DELETEExpressions: nil,
				REMOVEExpressions: nil,
				SETExpressions:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUpdateExpression(tt.args.updateExpression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseUpdateExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
