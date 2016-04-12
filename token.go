package regnet_lex_parse

type tokenType int

const (
	IILEGAL tokenType = iota
	KEY
	EQUALS
	PERCENT
	START_BRACE
	END_BRACE
	REG_TEXT
	WHITE_SPACE
)

const eof = rune(0)
