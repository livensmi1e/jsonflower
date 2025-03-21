package lexer

import (
	"strings"
	"unicode"
)

type LexFn func(*Lexer) LexFn

func LexValue(l *Lexer) LexFn {
	for {
		switch {
		case strings.HasPrefix(l.InputCurrentToEnd(), BEGIN_ARRAY):
			return LexBeginArray
		case strings.HasPrefix(l.InputCurrentToEnd(), BEGIN_OBJECT):
			return LexBeginObject
		case strings.HasPrefix(l.InputCurrentToEnd(), TRUE_LITERAL):
			return LexTrueLiteral
		case strings.HasPrefix(l.InputCurrentToEnd(), FALSE_LITERAL):
			return LexFalseLiteral
		case strings.HasPrefix(l.InputCurrentToEnd(), NULL_LITERAL):
			return LexNullLiteral
		case strings.HasPrefix(l.InputCurrentToEnd(), VALUE_SEPARATOR):
			return LexValueSeparator
		case strings.HasPrefix(l.InputCurrentToEnd(), NAME_SEPARATOR):
			return LexNameSeparator
		}
		switch r := l.Next(); {
		case r == EOF:
			return l.Errorf("unexpected end of file")
		case unicode.IsSpace(r):
			l.Ignore()
		case r == '-' || ('0' <= r && r <= '9'):
			l.Backup()
			return LexNumber
		case r == '"':
			return LexString
		default:
			return l.Errorf("unexpected character: %q", r)
		}
	}
}

func LexBeginArray(l *Lexer) LexFn {
	l.Position += len(BEGIN_ARRAY)
	l.Emit(TOKEN_BEGIN_ARRAY)
	return LexValue
}

func LexEndArray(l *Lexer) LexFn {
	l.Position += len(END_ARRAY)
	l.Emit(TOKEN_END_ARRAY)
	return LexValue
}

func LexValueSeparator(l *Lexer) LexFn {
	l.Position += len(VALUE_SEPARATOR)
	l.Emit(TOKEN_VALUE_SEPARATOR)
	return LexValue
}

func LexBeginObject(l *Lexer) LexFn {
	l.Position += len(BEGIN_OBJECT)
	l.Emit(TOKEN_BEGIN_OBJECT)
	return LexValue
}

func LexEndObject(l *Lexer) LexFn {
	l.Position += len(END_OBJECT)
	l.Emit(TOKEN_END_OBJECT)
	return LexValue
}

func LexNameSeparator(l *Lexer) LexFn {
	l.Position += len(NAME_SEPARATOR)
	l.Emit(TOKEN_NAME_SEPARATOR)
	return LexValue
}

func LexTrueLiteral(l *Lexer) LexFn {
	l.Position += len(TRUE_LITERAL)
	l.Emit(TOKEN_TRUE_LITERAL)
	return LexValue
}

func LexFalseLiteral(l *Lexer) LexFn {
	l.Position += len(FALSE_LITERAL)
	l.Emit(TOKEN_FALSE_LITERAL)
	return LexValue
}

func LexNullLiteral(l *Lexer) LexFn {
	l.Position += len(NULL_LITERAL)
	l.Emit(TOKEN_NULL_LITERAL)
	return LexValue
}

func LexNumber(l *Lexer) LexFn {
	l.Accept("-")
	digit1_9 := "123456789"
	digit0_9 := "012345678"
	if l.Accept("0") {

	} else if l.Accept(digit1_9) {
		l.AcceptRun(digit0_9)
	} else {
		return l.Errorf("bad number syntax %q", l.Input[l.Start:l.Position])
	}
	if l.Accept(".") {
		if !l.Accept(digit0_9) {
			return l.Errorf("bad number syntax %q", l.Input[l.Start:l.Position])
		}
		l.AcceptRun(digit0_9)
	}
	if l.Accept("Ee") {
		l.Accept("+-")
		if !l.Accept(digit0_9) {
			return l.Errorf("bad number syntax %q", l.Input[l.Start:l.Position])
		}
		l.AcceptRun(digit0_9)
	}
	l.Emit(TOKEN_NUMBER)
	return LexValue
}

// TODO: Need to handle \uXXXX
func LexString(l *Lexer) LexFn {
	for {
		switch l.Next() {
		case '"':
			l.Emit(TOKEN_STRING)
			return LexValue
		case '\\':
			if r := l.Next(); r == '/' || r == 'b' || r == 'f' || r == 'n' || r == 'r' || r == 't' {
				break
			}
			fallthrough
		case EOF, '\n':
			return l.Errorf("unterminated string")
		}
	}
}
