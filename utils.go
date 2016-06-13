package regnet_lex_parse

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isAlphabet(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isValidCharacters(ch rune) bool {
	return isAlphabet(ch) || isDigit(ch)
}
