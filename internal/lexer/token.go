package lexer

import "fmt"

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF
	TOKEN_BEGIN_ARRAY
	TOKEN_END_ARRAY
	TOKEN_BEGIN_OBJECT
	TOKEN_END_OBJECT
	TOKEN_NAME_SEPARATOR
	TOKEN_VALUE_SEPARATOR
	TOKEN_NUMBER
	TOKEN_STRING
	TOKEN_TRUE_LITERAL
	TOKEN_FALSE_LITERAL
	TOKEN_NULL_LITERAL
)

const (
	EOF rune = 0
)

const (
	BEGIN_ARRAY     = "["
	END_ARRAY       = "]"
	BEGIN_OBJECT    = "{"
	END_OBJECT      = "}"
	NAME_SEPARATOR  = ":"
	VALUE_SEPARATOR = ","
	TRUE_LITERAL    = "true"
	FALSE_LITERAL   = "false"
	NULL_LITERAL    = "null"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case TOKEN_EOF:
		return "EOF"
	case TOKEN_ERROR:
		return t.Value
	}
	if len(t.Value) > 10 {
		return fmt.Sprintf("%.10q...", t.Value)
	}
	return fmt.Sprintf("Type: %d - Value: %q", t.Type, t.Value)
}
