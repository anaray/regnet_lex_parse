package regnet_lex_parse

import (
	"fmt"
	"unicode/utf8"
)

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	if len(t.val) > 10 {
		return fmt.Sprintf("%10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	state  stateFn
	tokens chan token
}

func lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token),
	}
	go l.run()
	return l
}

type stateFn func(*lexer) stateFn

func (l *lexer) run() {
	for l.state = lexKey; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(t tokenType) {
	//fmt.Println("inside emit...")
	l.tokens <- token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}
