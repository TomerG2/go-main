package models

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	exp := "https://api.github.com/search/code?q=class+in%3Afile+language%3Ajava+repo%3Ascala%2Fscala"
	lines := []string{
		"https://api.github.com/search/code?q=class in:file language:java repo:scala/scala",
	}
	res := ParseLines(lines)
	if len(res) != 1 {
		t.Fatalf("Wrong number of urls")
	}
	if res[0] != exp {
		t.Fatalf("URLs parsed incorrectly")
	}
}
