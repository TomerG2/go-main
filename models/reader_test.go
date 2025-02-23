package models

import (
	"testing"
	"tomerg2/go-main/conf"
)

func TestParseFile(t *testing.T) {
	exp := "https://api.github.com/search/code?q=class in:file language:java repo:scala/scala"
	res := ReadFile(conf.QueryList)
	if len(res) != 5 {
		t.Fatalf("Missing expected results")
	}
	if res[0] != exp {
		t.Fatalf("Wrong lines")
	}
}
