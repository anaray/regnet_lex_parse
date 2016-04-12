package regnet_lex_parse

import (
	"strings"
)

func lexKey(l *lexer) stateFn {
	l.pos = strings.Index(l.input, "=")
	l.emit(KEY)
	return nil
}
