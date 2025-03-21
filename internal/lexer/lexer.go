package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Name     string
	Input    string
	State    LexFn
	Start    int
	Position int
	Width    int
	Tokens   chan Token
}

func (l *Lexer) Emit(t TokenType) {
	l.Tokens <- Token{Type: t, Value: l.Input[l.Start:l.Position]}
}

func (l *Lexer) Run() {
	for state := LexValue; state != nil; {
		state = state(l)
	}
}

func (l *Lexer) Accept(valid string) bool {
	if strings.ContainsRune(valid, l.Next()) {
		return true
	}
	l.Backup()
	return false
}

func (l *Lexer) AcceptRun(valid string) {
	for strings.ContainsRune(valid, l.Next()) {
	}
	l.Backup()
}

func (l *Lexer) Ignore() {
	l.Start = l.Position
}

func (l *Lexer) Backup() {
	l.Position -= l.Width
}

func (l *Lexer) Peek() rune {
	rune := l.Next()
	l.Backup()
	return rune
}

func (l *Lexer) Next() rune {
	if l.IsEOF() {
		l.Width = 0
		return EOF
	}
	result, Width := utf8.DecodeRuneInString(l.InputCurrentToEnd())
	l.Width = Width
	l.Position += Width
	return result
}

func (l *Lexer) InputCurrentToEnd() string {
	return l.Input[l.Position:]
}

func (l *Lexer) IsEOF() bool {
	return l.Position >= len(l.Input)
}

func (l *Lexer) SkipWhiteSpace() {
	for {
		r := l.Next()
		if !unicode.IsSpace(r) {
			l.Dec()
			break
		}
		if r == EOF {
			l.Emit(TOKEN_EOF)
			break
		}
	}
}

func (l *Lexer) Inc() {
	l.Position++
}

func (l *Lexer) Dec() {
	l.Position--
}

// TODO: Accept & AcceptRun
