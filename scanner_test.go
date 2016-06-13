package regnet_lex_parse

import (
	"testing"
  "strings"
  "fmt"
)

var s *Scanner

func TestScanner(t *testing.T){
  s = NewScanner(strings.NewReader("NAME=%{NAME_DEF}"))
}
