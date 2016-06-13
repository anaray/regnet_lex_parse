package regnet_lex_parse

type TokenType int

const (
	IILEGAL TokenType = iota
	KEY
	EQUALS
	PERCENT
	START_BRACE
	END_BRACE
	REG_TEXT
	WHITE_SPACE
	EOF
	SOME_TEXT_UNMARKED
)

const eof = rune(0)

type Position int

type Token struct {
  Type TokenType
	Pos Position
  Text string
}

func (t Token) String() string {
		return t.Text
}
