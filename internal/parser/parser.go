package parser

import "github.com/livensmi1e/jsonflower/internal/lexer"

type Parser struct {
	lex      *lexer.Lexer
	curToken lexer.Token
	error    error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l}
	return p
}
func (p *Parser) Parse() Value {
	go p.lex.Run()
	p.nextToken()
	return p.parseValue()
}

func (p *Parser) nextToken() {
	p.curToken = p.lex.NextToken()
}

func (p *Parser) hasError() bool {
	return p.error != nil
}
