package lexer

import "fmt"

const (
	LEXER_ERROR_UNEXPECTED_EOF     = "Unexpected end of file"
	LEXER_ERROR_MISSING_END_ARRAY  = "Missing an end array square bracket"
	LEXER_ERROR_MISSING_END_OBJECT = "Missing an end object curly bracket"
	// TODO: Define more errors if any
)

func (l *Lexer) Errorf(format string, args ...any) lexFn {
	l.tokens <- Token{
		TOKEN_ERROR,
		fmt.Sprintf(format, args...),
	}
	return nil
}
