package dynamock

import (
	"reflect"
	"testing"
)

func Test_mustExtractPathValueExpressions(t *testing.T) {
	type args struct {
		operation operation
		expr      string
	}
	tests := []struct {
		name string
		args args
		want []pathValueExpression
	}{
		{
			"no matching expression produces empty result",
			args{ADD, "SET foo = 3"},
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
		{
			"SET with single pair captures correct pair",
			args{SET, "SET foobar = 3"},
			[]pathValueExpression{{"foobar", "3"}},
		},
		{
			"SET with multiple pairs captures the pairs",
			args{SET, "SET foobar= 3, bazdog     =7, chicken=8    "},
			[]pathValueExpression{
				{"foobar", "3"},
				{"bazdog", "7"},
				{"chicken", "8"},
			},
		},
		{
			"SET with single pair including function captures correct pair",
			args{SET, "SET foobar = list_append(:vals, #ri)"},
			[]pathValueExpression{{"foobar", "list_append(:vals, #ri)"}},
		},
		{
			"SET with multiple pairs including function captures correct pair",
			args{SET, "SET dog = :food, foobar = list_append(:vals, #ri)"},
			[]pathValueExpression{
				{"dog", ":food"},
				{"foobar", "list_append(:vals, #ri)"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mustExtractPathValueExpressions(tt.args.operation, tt.args.expr)
			if len(got) != len(tt.want) {
				t.Errorf("mustExtractPathValueExpressions() len = %v, wanted len %v", len(got), len(tt.want))
				return
			}
			for idx, pair := range got {
				if !reflect.DeepEqual(pair, tt.want[idx]) {
					t.Errorf("mustExtractPathValueExpressions() %vth item got: %v, want: %v", idx, pair, tt.want[idx])
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
		{
			"SET expression only has only Set expressions",
			args{"SET foobar = 5, cat = list_append(:vals, #ri)"},
			parsedUpdateExpression{
				ADDExpressions:    nil,
				DELETEExpressions: nil,
				REMOVEExpressions: nil,
				SETExpressions: []pathValueExpression{
					{"foobar", "5"},
					{"cat", "list_append(:vals, #ri)"},
				},
			},
		},
		{
			"REMOVE expression only has only Remove expressions",
			args{"REMOVE RelatedItems[1], RelatedItems[2]"},
			parsedUpdateExpression{
				ADDExpressions:    nil,
				DELETEExpressions: nil,
				REMOVEExpressions: []pathExpression{
					{"RelatedItems[1]"},
					{"RelatedItems[2]"},
				},
				SETExpressions: nil,
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

func Test_extractRemovePath(t *testing.T) {
	type args struct {
		removeExpr string
	}
	tests := []struct {
		name string
		args args
		want []pathExpression
	}{
		{
			"extracting a single Remove path works correctly",
			args{"REMOVE RelatedItems[1]"},
			[]pathExpression{{"RelatedItems[1]"}},
		},
		{
			"extracting multiple Remove paths works correctly",
			args{"REMOVE RelatedItems[1], RelatedItems[2]"},
			[]pathExpression{
				{"RelatedItems[1]"},
				{"RelatedItems[2]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractRemovePath(tt.args.removeExpr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractRemovePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsedUpdateExpression_CheckIsEquivalentTo(t *testing.T) {
	type args struct {
		p     string
		other string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"equivalent expressions do not return an error",
			args{
				"ADD foobar 5, dog 7 ",
				"ADD dog 7, foobar 5",
			},
			false,
		},
		{
			"non-equivalent expressions return an error",
			args{
				"ADD foobar 5, dog 7 ",
				"ADD cat 7, foobar 5",
			},
			true,
		},
		{
			"non-equivalent expressions with equivalent paths return an error",
			args{
				"ADD foobar 5, dog 7 ",
				"ADD dog 7, foobar 99",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pExpr := parseUpdateExpression(tt.args.p)
			otherExpr := parseUpdateExpression(tt.args.other)
			if err := pExpr.CheckIsEquivalentTo(&otherExpr); (err != nil) != tt.wantErr {
				t.Errorf("CheckIsEquivalentTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
