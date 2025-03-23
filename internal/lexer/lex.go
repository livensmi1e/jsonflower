package lexer

import (
	"strings"
	"unicode"
)

type lexFn func(*Lexer) lexFn

func lexValue(l *Lexer) lexFn {
	for {
		switch {
		case strings.HasPrefix(l.posToEndInput(), BEGIN_ARRAY):
			return lexBeginArray
		case strings.HasPrefix(l.posToEndInput(), BEGIN_OBJECT):
			return lexBeginObject
		case strings.HasPrefix(l.posToEndInput(), TRUE_LITERAL):
			return lexTrueLiteral
		case strings.HasPrefix(l.posToEndInput(), FALSE_LITERAL):
			return lexFalseLiteral
		case strings.HasPrefix(l.posToEndInput(), NULL_LITERAL):
			return lexNullLiteral
		case strings.HasPrefix(l.posToEndInput(), VALUE_SEPARATOR):
			return lexValueSeparator
		case strings.HasPrefix(l.posToEndInput(), NAME_SEPARATOR):
			return lexNameSeparator
		case strings.HasPrefix(l.posToEndInput(), END_ARRAY):
			return lexEndArray
		case strings.HasPrefix(l.posToEndInput(), END_OBJECT):
			return lexEndObject
		}
		switch r := l.next(); {
		case r == EOF:
			l.emit(TOKEN_EOF)
			return nil
		case unicode.IsSpace(r):
			l.ignore()
		case r == '-' || ('0' <= r && r <= '9'):
			l.backup()
			return lexNumber
		case r == '"':
			return lexString
		default:
			return l.errorf("unexpected character: %q", r)
		}
	}
}

func lexBeginArray(l *Lexer) lexFn {
	l.pos += len(BEGIN_ARRAY)
	l.emit(TOKEN_BEGIN_ARRAY)
	return lexValue
}

func lexEndArray(l *Lexer) lexFn {
	l.pos += len(END_ARRAY)
	l.emit(TOKEN_END_ARRAY)
	return lexValue
}

func lexValueSeparator(l *Lexer) lexFn {
	l.pos += len(VALUE_SEPARATOR)
	l.emit(TOKEN_VALUE_SEPARATOR)
	return lexValue
}

func lexBeginObject(l *Lexer) lexFn {
	l.pos += len(BEGIN_OBJECT)
	l.emit(TOKEN_BEGIN_OBJECT)
	return lexValue
}

func lexEndObject(l *Lexer) lexFn {
	l.pos += len(END_OBJECT)
	l.emit(TOKEN_END_OBJECT)
	return lexValue
}

func lexNameSeparator(l *Lexer) lexFn {
	l.pos += len(NAME_SEPARATOR)
	l.emit(TOKEN_NAME_SEPARATOR)
	return lexValue
}

func lexTrueLiteral(l *Lexer) lexFn {
	l.pos += len(TRUE_LITERAL)
	l.emit(TOKEN_TRUE_LITERAL)
	return lexValue
}

func lexFalseLiteral(l *Lexer) lexFn {
	l.pos += len(FALSE_LITERAL)
	l.emit(TOKEN_FALSE_LITERAL)
	return lexValue
}

func lexNullLiteral(l *Lexer) lexFn {
	l.pos += len(NULL_LITERAL)
	l.emit(TOKEN_NULL_LITERAL)
	return lexValue
}

func lexNumber(l *Lexer) lexFn {
	l.accept("-")
	digit1_9 := "123456789"
	digit0_9 := "012345678"
	if l.accept("0") {

	} else if l.accept(digit1_9) {
		l.acceptRun(digit0_9)
	} else {
		return l.errorf("bad number syntax %q", l.input[l.start:l.pos])
	}
	if l.accept(".") {
		if !l.accept(digit0_9) {
			return l.errorf("bad number syntax %q", l.input[l.start:l.pos])
		}
		l.acceptRun(digit0_9)
	}
	if l.accept("Ee") {
		l.accept("+-")
		if !l.accept(digit0_9) {
			return l.errorf("bad number syntax %q", l.input[l.start:l.pos])
		}
		l.acceptRun(digit0_9)
	}
	l.emit(TOKEN_NUMBER)
	return lexValue
}

// TODO: Need to handle \uXXXX
func lexString(l *Lexer) lexFn {
	for {
		switch l.next() {
		case '"':
			l.emit(TOKEN_STRING)
			return lexValue
		case '\\':
			if r := l.next(); r == '/' || r == 'b' || r == 'f' || r == 'n' || r == 'r' || r == 't' {
				break
			}
			fallthrough
		case EOF, '\n':
			return l.errorf("unterminated string")
		}
	}
}
