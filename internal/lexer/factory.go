package lexer

func New(name, input string) *Lexer {
	l := &Lexer{
		Name:   name,
		Input:  input,
		State:  LexValue,
		Tokens: make(chan Token, 3),
	}
	return l
}
