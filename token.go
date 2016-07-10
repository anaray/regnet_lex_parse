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
	Pos  Position
	Text string
}

func (t Token) String() string {
	return t.Text
}

func (typ TokenType) String() string {
	switch typ {
	case IILEGAL:
		return "IILEGAL"
	case KEY:
		return "KEY"
	case EQUALS:
		return "EQUALS"
	case PERCENT:
		return "PERCENT"
	case START_BRACE:
		return "START_BRACE"
	case END_BRACE:
		return "END_BRACE"
	case REG_TEXT:
		return "REG_TEXT"
	case WHITE_SPACE:
		return "WHITE_SPACE"
	case EOF:
		return "EOF"
	default:
		panic("unrecognized token")
	}
}
