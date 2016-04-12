package regnet_lex_parse

import (
	"testing"
)

var l *lexer

func Test_Lex(t *testing.T) {
	l = lex("NAME=%{NAME_DEF}")
	if l == nil {
		t.Fatalf("lexer and tokens channel expected")
	}
}

func TestLexKeyEmit(t *testing.T) {
	l := lex("NAME=%{NAME_DEF}")

	for token := range l.tokens {
		if token.val != "NAME" {
			t.Fatalf("expected Key=NAME, but received %s", token.val)
		}
	}
}
