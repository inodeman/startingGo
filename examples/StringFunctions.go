package examples

import ss "strings"
import "fmt"

var p = fmt.Println

func StringFunctions() {
	p("** Examples String Functions**")
	p("Contains:   ", ss.Contains("test", "es"))
	p("Count:      ", ss.Count("test", "t"))
	p("HasPrefix:  ", ss.HasPrefix("test", "te"))
	p("HasSuffix:  ", ss.HasSuffix("test", "st"))
	p("Index:      ", ss.Index("test", "e"))
	p("Join:       ", ss.Join([]string{"a", "b"}, "-"))
	p("Repeat:     ", ss.Repeat("a", 5))
	p("Replace:    ", ss.Replace("foo", "o", "0", -1))
	p("Replace:    ", ss.Replace("foo", "o", "0", 1))
	p("Split:      ", ss.Split("a-b-c-d-e", "-"))
	p("ToLower:    ", ss.ToLower("TEST"))
	p("ToUpper:    ", ss.ToUpper("test"))
	p()
	p("Len:        ", len("hello"))
	p("Char:       ", "hello"[1])
}
