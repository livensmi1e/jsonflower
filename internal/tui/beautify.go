package tui

import (
	"github.com/livensmi1e/jsonflower/internal/beautify"
	"github.com/livensmi1e/jsonflower/internal/lexer"
	"github.com/livensmi1e/jsonflower/internal/parser"
)

func beautifyJSON(input string) string {
	l := lexer.New("JSON", input)
	p := parser.New(l)
	ast := p.Parse()
	if p.Err() != "" {
		return "❌ Invalid JSON"
	}
	return beautify.BeautifyAST(ast)
}
