package regnet_lex_parse

import (
	"strings"
	"testing"
)

var s *Scanner

func TestScanner(t *testing.T) {
	s = NewScanner(strings.NewReader("NAME=%{NAME_DEF}"))
	token := s.Scan()
	if token.Type != KEY {
		t.Error("expected token type KEY")
	}

	if token.Text != "NAME" {
		t.Error("expected token text NAME")
	}

	token = s.Scan()
	if token.Type != EQUALS {
		t.Error("expected token type EQUALS")
	}

	if token.Text != "=" {
		t.Error("expected token text EQUALS")
	}

	s = NewScanner(strings.NewReader(" "))
	token = s.Scan()

	if token.Type != EOF {
		t.Error("expected token type EOF")
	}

	s = NewScanner(strings.NewReader("NAME=%{NAME_DEF}"))
	ch := s.Start()
	count := 0
	for range ch {
		count++
	}
	if count != 6 {
		t.Error("expected 6 tokens")
	}

	s = NewScanner(strings.NewReader("NAME=%{FIRST_NAME}%{LAST_NAME"))
	ch = s.Start()
	count = 0
	for range ch {
		count++
	}
	if count != 9 {
		t.Error("expected 8 tokens")
	}

}
