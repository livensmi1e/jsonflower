package lexer

import (
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	name        string
	input       string
	initalState lexFn
	start       int
	pos         int
	width       int
	tokens      chan Token
}

func (l *Lexer) Run() {
	defer close(l.tokens)
	for state := l.initalState; state != nil; {
		state = state(l)
	}
}

func (l *Lexer) NextToken() Token {
	return <-l.tokens
}

func (l *Lexer) emit(t TokenType) {
	l.tokens <- Token{Type: t, Lexeme: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) next() rune {
	if l.isEOF() {
		l.width = 0
		return EOF
	}
	result, width := utf8.DecodeRuneInString(l.posToEndInput())
	l.width = width
	l.pos += width
	return result
}

func (l *Lexer) posToEndInput() string {
	return l.input[l.pos:]
}

func (l *Lexer) isEOF() bool {
	return l.pos >= len(l.input)
}
