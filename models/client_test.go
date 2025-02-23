package models

import "testing"

func TestAddPagination(t *testing.T) {
	url := "https://api.github.com/search/code?q=class+in%3Afile+language%3Ajava+repo%3Ascala%2Fscala"
	exp := "https://api.github.com/search/code?page=1&per_page=10&q=class+in%3Afile+language%3Ajava+repo%3Ascala%2Fscala"
	res := AddPagination(url, 1, 10)
	if res != exp {
		t.Fatalf("Params not added")
	}
}
