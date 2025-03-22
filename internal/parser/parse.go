package parser

import (
	"fmt"
	"strconv"

	"github.com/livensmi1e/jsonflower/internal/lexer"
)

func (p *Parser) parseValue() Value {
	if p.hasError() {
		return nil
	}
	switch p.curToken.Type {
	case lexer.TOKEN_BEGIN_OBJECT:
		p.nextToken()
		return p.parseObject()
	case lexer.TOKEN_BEGIN_ARRAY:
		p.nextToken()
		return p.parseArray()
	case lexer.TOKEN_STRING:
		s := String{Literal: p.curToken.Value}
		p.nextToken()
		return s
	case lexer.TOKEN_NUMBER:
		v, _ := strconv.ParseFloat(p.curToken.Value, 64)
		n := Number{Value: v}
		p.nextToken()
		return n
	case lexer.TOKEN_TRUE_LITERAL, lexer.TOKEN_FALSE_LITERAL:
		v, _ := strconv.ParseBool(p.curToken.Value)
		b := Boolean{Value: v}
		p.nextToken()
		return b
	case lexer.TOKEN_NULL_LITERAL:
		p.nextToken()
		return Null{}
	default:
		p.error = fmt.Errorf("unexpected token: %s", p.curToken.String())
		return nil
	}
}

func (p *Parser) parseObject() Value {
	if p.hasError() {
		return nil
	}
	o := &Object{KeyValue: make(map[string]Value)}
	if p.curToken.Type == lexer.TOKEN_END_OBJECT {
		p.nextToken()
		return o
	}
	for !p.hasError() {
		key, value := p.parseMember()
		if p.hasError() {
			return nil
		}
		o.KeyValue[key] = value
		if p.curToken.Type == lexer.TOKEN_END_OBJECT {
			p.nextToken()
			break
		}
		if p.curToken.Type != lexer.TOKEN_VALUE_SEPARATOR {
			p.error = fmt.Errorf("expected ',', got %s", p.curToken.String())
			return nil
		}
		p.nextToken()
	}
	return o
}

func (p *Parser) parseMember() (string, Value) {
	if p.hasError() {
		return "", nil
	}
	if p.curToken.Type != lexer.TOKEN_STRING {
		p.error = fmt.Errorf("expected string key, got %s", p.curToken.String())
		return "", nil
	}
	key := p.curToken.Value
	p.nextToken()
	if p.curToken.Type != lexer.TOKEN_NAME_SEPARATOR {
		p.error = fmt.Errorf("expected ':', got %s", p.curToken.String())
		return "", nil
	}
	p.nextToken()
	value := p.parseValue()
	return key, value
}

func (p *Parser) parseArray() Value {
	if p.hasError() {
		return nil
	}
	a := &Array{}
	if p.curToken.Type == lexer.TOKEN_END_ARRAY {
		p.nextToken()
		return a
	}
	for !p.hasError() {
		value := p.parseValue()
		if p.hasError() {
			return nil
		}
		a.Elements = append(a.Elements, value)
		if p.curToken.Type == lexer.TOKEN_END_ARRAY {
			p.nextToken()
			break
		}
		if p.curToken.Type != lexer.TOKEN_VALUE_SEPARATOR {
			p.error = fmt.Errorf("expected ',', got %s", p.curToken.String())
			return nil
		}
		p.nextToken()
	}
	return a
}
