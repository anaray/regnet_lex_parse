package regnet_lex_parse

import (
	"fmt"
	"strings"
	"testing"
)

var s *Scanner

func TestScanner(t *testing.T) {
	s = NewScanner(strings.NewReader("NAME=%{NAME_DEF}"))
}
