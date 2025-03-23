package lexer

func New(name, input string) *Lexer {
	l := &Lexer{
		name:        name,
		input:       input,
		initalState: LexValue,
		tokens:      make(chan Token, 3),
	}
	return l
}
