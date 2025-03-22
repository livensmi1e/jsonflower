package lexer

import "testing"

func TestLexer(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []Token
	}{
		{
			name:  "Simple JSON",
			input: `{"key": 123, "valid": true, "pi": 3.14, "msg": "hello"}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"key\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "123"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"valid\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_TRUE_LITERAL, Value: TRUE_LITERAL},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"pi\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "3.14"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"msg\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"hello\""},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
		{
			name:  "Empty JSON",
			input: `{}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
		{
			name:  "Nested JSON",
			input: `{"user": {"name": "Alice", "age": 25}}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"user\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"name\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"Alice\""},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"age\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "25"},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
		{
			name:  "Array JSON",
			input: `{"numbers": [1, 2, 3]}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"numbers\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Value: BEGIN_ARRAY},
				{Type: TOKEN_NUMBER, Value: "1"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "2"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "3"},
				{Type: TOKEN_END_ARRAY, Value: END_ARRAY},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
		{
			name:  "Whitespace Handling",
			input: `{   "key"  :  "value"   }`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"key\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"value\""},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
		{
			name: "JSON with nested objects and arrays",
			input: `{
				"numbers": [1, 2, 3.5],
				"nested": {
					"valid": true,
					"emptyArray": [],
					"nullValue": null
				}
			}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"numbers\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Value: BEGIN_ARRAY},
				{Type: TOKEN_NUMBER, Value: "1"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "2"},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Value: "3.5"},
				{Type: TOKEN_END_ARRAY, Value: END_ARRAY},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"nested\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_OBJECT, Value: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Value: "\"valid\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_TRUE_LITERAL, Value: TRUE_LITERAL},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"emptyArray\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Value: BEGIN_ARRAY},
				{Type: TOKEN_END_ARRAY, Value: END_ARRAY},
				{Type: TOKEN_VALUE_SEPARATOR, Value: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Value: "\"nullValue\""},
				{Type: TOKEN_NAME_SEPARATOR, Value: NAME_SEPARATOR},
				{Type: TOKEN_NULL_LITERAL, Value: NULL_LITERAL},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_END_OBJECT, Value: END_OBJECT},
				{Type: TOKEN_EOF, Value: "EOF"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			l := New("test", tc.input)
			go l.Run()
			for i, expected := range tc.expectedTokens {
				token := l.NextToken()
				if token.Type != expected.Type {
					t.Errorf("Token %d: expected %s, got %s", i, expected.String(), token.String())
				}
			}
			select {
			case token, ok := <-l.Tokens:
				if ok {
					t.Errorf("Unexpected extra token: %s", token.String())
				}
			default:
			}
		})
	}
}
