package parser

// TODO: May I need some method?
type Value interface{}

type Object struct {
	KeyValue map[string]Value
}

type Array struct {
	Elements []Value
}

type Number struct {
	Value float64
}

type String struct {
	Literal string
}

type Boolean struct {
	Value bool
}

type Null struct{}
