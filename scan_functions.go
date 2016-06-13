package regnet_lex_parse

import (
	"bytes"
)

func (s *Scanner) scanWhiteSpace() (token Token) {
	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			return Token{Type: EOF, Text: string(ch)}
			//break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return Token{Type: WHITE_SPACE, Text: buf.String()}
}

func (s *Scanner) scanValidCharacters() (token Token) {
	var buf bytes.Buffer
	for {
		if ch := s.read(); ch == eof {
			return Token{Type: EOF, Text: string(ch)}
		} else if isAlphabet(ch) {
			buf.WriteRune(ch)
		} else if isDigit(ch) {
			buf.WriteRune(ch)
		} else if isWhitespace(ch) {
			s.unread()
			break
		} else if ch == '=' {
			s.unread()
			return Token{Type: KEY, Text: buf.String()}
		} else if ch == '}' {
			s.unread()
			return Token{Type: REG_TEXT, Text: buf.String()}
		}
	}
	return Token{Type: SOME_TEXT_UNMARKED, Text: buf.String()}
}
