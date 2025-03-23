package lexer

import (
	"strings"
	"unicode"
)

type lexFn func(*Lexer) lexFn

func LexValue(l *Lexer) lexFn {
	for {
		switch {
		case strings.HasPrefix(l.posToEndInput(), BEGIN_ARRAY):
			return LexBeginArray
		case strings.HasPrefix(l.posToEndInput(), BEGIN_OBJECT):
			return LexBeginObject
		case strings.HasPrefix(l.posToEndInput(), TRUE_LITERAL):
			return LexTrueLiteral
		case strings.HasPrefix(l.posToEndInput(), FALSE_LITERAL):
			return LexFalseLiteral
		case strings.HasPrefix(l.posToEndInput(), NULL_LITERAL):
			return LexNullLiteral
		case strings.HasPrefix(l.posToEndInput(), VALUE_SEPARATOR):
			return LexValueSeparator
		case strings.HasPrefix(l.posToEndInput(), NAME_SEPARATOR):
			return LexNameSeparator
		case strings.HasPrefix(l.posToEndInput(), END_ARRAY):
			return LexEndArray
		case strings.HasPrefix(l.posToEndInput(), END_OBJECT):
			return LexEndObject
		}
		switch r := l.next(); {
		case r == EOF:
			l.emit(TOKEN_EOF)
			return nil
		case unicode.IsSpace(r):
			l.ignore()
		case r == '-' || ('0' <= r && r <= '9'):
			l.backup()
			return LexNumber
		case r == '"':
			return LexString
		default:
			return l.Errorf("unexpected character: %q", r)
		}
	}
}

func LexBeginArray(l *Lexer) lexFn {
	l.pos += len(BEGIN_ARRAY)
	l.emit(TOKEN_BEGIN_ARRAY)
	return LexValue
}

func LexEndArray(l *Lexer) lexFn {
	l.pos += len(END_ARRAY)
	l.emit(TOKEN_END_ARRAY)
	return LexValue
}

func LexValueSeparator(l *Lexer) lexFn {
	l.pos += len(VALUE_SEPARATOR)
	l.emit(TOKEN_VALUE_SEPARATOR)
	return LexValue
}

func LexBeginObject(l *Lexer) lexFn {
	l.pos += len(BEGIN_OBJECT)
	l.emit(TOKEN_BEGIN_OBJECT)
	return LexValue
}

func LexEndObject(l *Lexer) lexFn {
	l.pos += len(END_OBJECT)
	l.emit(TOKEN_END_OBJECT)
	return LexValue
}

func LexNameSeparator(l *Lexer) lexFn {
	l.pos += len(NAME_SEPARATOR)
	l.emit(TOKEN_NAME_SEPARATOR)
	return LexValue
}

func LexTrueLiteral(l *Lexer) lexFn {
	l.pos += len(TRUE_LITERAL)
	l.emit(TOKEN_TRUE_LITERAL)
	return LexValue
}

func LexFalseLiteral(l *Lexer) lexFn {
	l.pos += len(FALSE_LITERAL)
	l.emit(TOKEN_FALSE_LITERAL)
	return LexValue
}

func LexNullLiteral(l *Lexer) lexFn {
	l.pos += len(NULL_LITERAL)
	l.emit(TOKEN_NULL_LITERAL)
	return LexValue
}

func LexNumber(l *Lexer) lexFn {
	l.accept("-")
	digit1_9 := "123456789"
	digit0_9 := "012345678"
	if l.accept("0") {

	} else if l.accept(digit1_9) {
		l.acceptRun(digit0_9)
	} else {
		return l.Errorf("bad number syntax %q", l.input[l.start:l.pos])
	}
	if l.accept(".") {
		if !l.accept(digit0_9) {
			return l.Errorf("bad number syntax %q", l.input[l.start:l.pos])
		}
		l.acceptRun(digit0_9)
	}
	if l.accept("Ee") {
		l.accept("+-")
		if !l.accept(digit0_9) {
			return l.Errorf("bad number syntax %q", l.input[l.start:l.pos])
		}
		l.acceptRun(digit0_9)
	}
	l.emit(TOKEN_NUMBER)
	return LexValue
}

// TODO: Need to handle \uXXXX
func LexString(l *Lexer) lexFn {
	for {
		switch l.next() {
		case '"':
			l.emit(TOKEN_STRING)
			return LexValue
		case '\\':
			if r := l.next(); r == '/' || r == 'b' || r == 'f' || r == 'n' || r == 'r' || r == 't' {
				break
			}
			fallthrough
		case EOF, '\n':
			return l.Errorf("unterminated string")
		}
	}
}
