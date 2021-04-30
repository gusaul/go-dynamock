package dynamock

import (
	"reflect"
	"testing"
)

func Test_mustExtractPathValueExpressions(t *testing.T) {
	type args struct {
		operation operation
		addExpr   string
	}
	tests := []struct {
		name string
		args args
		want []pathValueExpression
	}{
		{
			"no ADD or DELETE in expression produces empty result",
			args{SET, "SET foo = 3"},
			[]pathValueExpression{},
		},
		{
			"ADD with single pair captures correct pair",
			args{ADD, "ADD foobar 3"},
			[]pathValueExpression{{"foobar", "3"}},
		},
		{
			"ADD with single pair and trailing comma captures correct pair",
			args{ADD, "ADD foobar 3 ,"},
			[]pathValueExpression{{"foobar", "3"}},
		},
		{
			"ADD with single pair and trailing whitespace captures correct pair",
			args{ADD, "ADD foobar 3   "},
			[]pathValueExpression{{"foobar", "3"}},
		},
		{
			"ADD with multiple pairs captures the pairs",
			args{ADD, "ADD foobar 3, bazdog 7, chicken 8"},
			[]pathValueExpression{
				{"foobar", "3"},
				{"bazdog", "7"},
				{"chicken", "8"},
			},
		},
		{
			"ADD with multiple pairs and curious whitespace",
			args{ADD, "ADD 	  foobar   	3  ,   bazdog 7   ,chicken     8"},
			[]pathValueExpression{
				{"foobar", "3"},
				{"bazdog", "7"},
				{"chicken", "8"},
			},
		},
		{
			"DELETE with single pair and trailing comma captures correct pair",
			args{DELETE, "DELETE foobar 3 ,"},
			[]pathValueExpression{{"foobar", "3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mustExtractPathValueExpressions(tt.args.operation, tt.args.addExpr)
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
				ADDExpressions: []pathValueExpression{
					{"foobar", "3"},
					{"dog", "76"},
				},
				DELETEExpressions: nil,
				REMOVEExpressions: nil,
				SETExpressions:    nil,
			},
		},
		{
			"DELETE expression only has only Delete expressions",
			args{"DELETE foobar 5, cat 6"},
			parsedUpdateExpression{
				ADDExpressions: nil,
				DELETEExpressions: []pathValueExpression{
					{"foobar", "5"},
					{"cat", "6"},
				},
				REMOVEExpressions: nil,
				SETExpressions:    nil,
			},
		},
		{
			"ADD and DELETE expressions are extracted correctly",
			args{"ADD abc 1, def 2 DELETE ghi 3, jkl 4"},
			parsedUpdateExpression{
				ADDExpressions: []pathValueExpression{
					{"abc", "1"},
					{"def", "2"},
				},
				DELETEExpressions: []pathValueExpression{
					{"ghi", "3"},
					{"jkl", "4"},
				},
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
