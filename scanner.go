package regnet_lex_parse

import (
	"bufio"
	"io"
)

type Scanner struct {
	reader  *bufio.Reader
	Channel chan Token
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		reader:  bufio.NewReader(r),
		Channel: make(chan Token),
	}
}

func (s *Scanner) Scan() (tok Token) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isValidCharacters(ch) {
		s.unread()
		return s.scanValidCharacters()
	}

	switch ch {
	case eof:
		return Token{Type: EOF, Text: ""}
	case '=':
		return Token{Type: EQUALS, Text: string(ch)}
	case '%':
		return Token{Type: PERCENT, Text: string(ch)}
	case '{':
		return Token{Type: START_BRACE, Text: string(ch)}
	case '}':
		return Token{Type: END_BRACE, Text: string(ch)}
	}
	return Token{Type: IILEGAL, Text: string(ch)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.reader.UnreadRune()
}

func (s *Scanner) peek() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return eof
	}
	s.reader.UnreadRune()
	return ch
}

func (s *Scanner) emit(t Token) {
	s.Channel <- t
}
