package tui

import (
	"github.com/livensmi1e/jsonflower/internal/beautify"
	"github.com/livensmi1e/jsonflower/internal/lexer"
	"github.com/livensmi1e/jsonflower/internal/parser"
	"github.com/livensmi1e/jsonflower/internal/transformer"
)

func beautifyJSON(input string) string {
	l := lexer.New("Beautify JSON", input)
	p := parser.New(l)
	ast := p.Parse()
	if p.Err() != "" {
		return "❌ Invalid JSON"
	}
	return beautify.BeautifyAST(ast)
}

func convertJSON2YAML(input string) string {
	l := lexer.New("JSON2YAML", input)
	p := parser.New(l)
	ast := p.Parse()
	if p.Err() != "" {
		return "❌ Invalid JSON"
	}
	return transformer.TransformJSON2YAML(ast)
}
