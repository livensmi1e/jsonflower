package parser

import (
	"testing"

	"github.com/livensmi1e/jsonflower/internal/lexer"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Value
	}{
		{"String", `"hello"`, String{Literal: "\"hello\""}},
		{"Number", `42`, Number{Value: 42}},
		{"Float", `3.14`, Number{Value: 3.14}},
		{"True", `true`, Boolean{Value: true}},
		{"False", `false`, Boolean{Value: false}},
		{"Null", `null`, Null{}},

		{"Empty Array", `[]`, &Array{}},
		{"Array of Numbers", `[1, 2, 3]`, &Array{Elements: []Value{
			Number{Value: 1}, Number{Value: 2}, Number{Value: 3},
		}}},
		{"Mixed Array", `["a", 1, true, null]`, &Array{Elements: []Value{
			String{Literal: "\"a\""}, Number{Value: 1}, Boolean{Value: true}, Null{},
		}}},

		{"Empty Object", `{}`, &Object{KeyValue: map[string]Value{}}},
		{"Simple Object", `{"key": "value"}`, &Object{KeyValue: map[string]Value{
			"\"key\"": String{Literal: "\"value\""},
		}}},
		{"Object with Numbers", `{"a": 1, "b": 2.5}`, &Object{KeyValue: map[string]Value{
			"\"a\"": Number{Value: 1}, "\"b\"": Number{Value: 2.5},
		}}},

		{"Invalid JSON", `{invalid}`, nil},
		{"Missing Comma", `{"a": 1 "b": 2}`, nil},
		{"Unclosed Array", `[1, 2, 3`, nil},
		{"Unclosed Object", `{"a": 1,`, nil},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.name, tc.input)
			p := New(l)
			parsed := p.Parse()
			if !equalValues(parsed, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, parsed)
			}
		})
	}
}

func equalValues(a, b Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	switch a := a.(type) {
	case String:
		b, ok := b.(String)
		return ok && a.Literal == b.Literal
	case Number:
		b, ok := b.(Number)
		return ok && a.Value == b.Value
	case Boolean:
		b, ok := b.(Boolean)
		return ok && a.Value == b.Value
	case Null:
		_, ok := b.(Null)
		return ok
	case *Array:
		b, ok := b.(*Array)
		if !ok || len(a.Elements) != len(b.Elements) {
			return false
		}
		for i := range a.Elements {
			if !equalValues(a.Elements[i], b.Elements[i]) {
				return false
			}
		}
		return true
	case *Object:
		b, ok := b.(*Object)
		if !ok || len(a.KeyValue) != len(b.KeyValue) {
			return false
		}
		for k, v := range a.KeyValue {
			bv, exists := b.KeyValue[k]
			if !exists || !equalValues(v, bv) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
