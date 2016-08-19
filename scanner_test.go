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

	token = <-ch
	if token.Type != KEY && token.Text != "NAME" {
		t.Error("expected token type KEY with value NAME")
	}

	token = <-ch
	if token.Type != EQUALS && token.Text != "=" {
		t.Error("expected token type EQUALS with value '=''")
	}

	token = <-ch
	if token.Type != PERCENT && token.Text != "%" {
		t.Error("expected token type PERCENT with value '%''")
	}

	token = <-ch
	if token.Type != START_BRACE && token.Text != "{" {
		t.Error("expected token type START_BRACE with value '{''")
	}

	token = <-ch
	if token.Type != REG_TEXT && token.Text != "FIRST_NAME" {
		t.Error("expected token type REG_TEXT with value 'FIRST_NAME''")
	}

	token = <-ch
	if token.Type != REG_TEXT && token.Text != "}" {
		t.Error("expected token type END_BRACE with value '}''")
	}

	token = <-ch
	if token.Type != PERCENT && token.Text != "%" {
		t.Error("expected token type PERCENT with value '%''")
	}

	token = <-ch
	if token.Type != START_BRACE && token.Text != "{" {
		t.Error("expected token type START_BRACE with value '{''")
	}

	token = <-ch
	if token.Type != REG_TEXT && token.Text != "LAST_NAME" {
		t.Error("expected token type REG_TEXT with value 'LAST_NAME''")
	}
}
