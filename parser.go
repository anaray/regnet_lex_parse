package regnet_lex_parse

import (
	"fmt"
	"io"
)

type Parser struct {
	scanner *Scanner
}

func NewParser(r io.Reader) *Parser {
	return &Parser{scanner: NewScanner(r)}
}

func (p *Parser) Start() {
	ch := p.scanner.Start()
	for token := range ch {
		fmt.Println(token)
	}
}
