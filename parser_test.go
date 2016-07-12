package regnet_lex_parse

import(
  "testing"
  "fmt"
  "strings"
)

func TestParser(t *testing.T) {
    p := NewParser(strings.NewReader("NAME=%{NAME_DEF}"))
    fmt.Println(p)
}
