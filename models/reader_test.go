package models

import (
	"testing"
	"tomerg2/go-main/conf"
)

func TestParseFile(t *testing.T) {
	res := ReadFile(conf.QueryList)
	if len(res) != 5 {
		t.Fatalf("Missing expected results")
	}
}
