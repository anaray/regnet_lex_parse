package regnet_lex_parse

import(
  "io"
)

type Parser struct {
  scanner *Scanner
}

func NewParser(r io.Reader) *Parser {
    return &Parser{scanner: NewScanner(r)}
}
