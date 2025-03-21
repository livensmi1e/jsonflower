package lexer

import "testing"

func TestLexer(t *testing.T) {
	input := `{"key": 123, "valid": true, "pi": 3.14, "msg": "hello"}`
	l := BeginLexing("test", input)
	go l.Run()
	expectedTokens := []Token{
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
	}
	for i, expected := range expectedTokens {
		token := l.NextToken()
		if token.Type != expected.Type || token.Value != expected.Value {
			t.Errorf("Token %d: expected %s, got %s", i, expected.String(), token.String())
		}
	}
}
