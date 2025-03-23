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
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"key\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "123"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"valid\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_TRUE_LITERAL, Lexeme: TRUE_LITERAL},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"pi\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "3.14"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"msg\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"hello\""},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
			},
		},
		{
			name:  "Empty JSON",
			input: `{}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
			},
		},
		{
			name:  "Nested JSON",
			input: `{"user": {"name": "Alice", "age": 25}}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"user\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"name\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"Alice\""},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"age\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "25"},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
			},
		},
		{
			name:  "Array JSON",
			input: `{"numbers": [1, 2, 3]}`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"numbers\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Lexeme: BEGIN_ARRAY},
				{Type: TOKEN_NUMBER, Lexeme: "1"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "2"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "3"},
				{Type: TOKEN_END_ARRAY, Lexeme: END_ARRAY},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
			},
		},
		{
			name:  "Whitespace Handling",
			input: `{   "key"  :  "value"   }`,
			expectedTokens: []Token{
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"key\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"value\""},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
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
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"numbers\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Lexeme: BEGIN_ARRAY},
				{Type: TOKEN_NUMBER, Lexeme: "1"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "2"},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_NUMBER, Lexeme: "3.5"},
				{Type: TOKEN_END_ARRAY, Lexeme: END_ARRAY},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"nested\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_OBJECT, Lexeme: BEGIN_OBJECT},
				{Type: TOKEN_STRING, Lexeme: "\"valid\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_TRUE_LITERAL, Lexeme: TRUE_LITERAL},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"emptyArray\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_BEGIN_ARRAY, Lexeme: BEGIN_ARRAY},
				{Type: TOKEN_END_ARRAY, Lexeme: END_ARRAY},
				{Type: TOKEN_VALUE_SEPARATOR, Lexeme: VALUE_SEPARATOR},
				{Type: TOKEN_STRING, Lexeme: "\"nullValue\""},
				{Type: TOKEN_NAME_SEPARATOR, Lexeme: NAME_SEPARATOR},
				{Type: TOKEN_NULL_LITERAL, Lexeme: NULL_LITERAL},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_END_OBJECT, Lexeme: END_OBJECT},
				{Type: TOKEN_EOF, Lexeme: "EOF"},
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
			case token, ok := <-l.tokens:
				if ok {
					t.Errorf("Unexpected extra token: %s", token.String())
				}
			default:
			}
		})
	}
}
