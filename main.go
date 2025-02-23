package main

import (
	"fmt"
	"tomerg2/go-main/conf"
	"tomerg2/go-main/models"
)

func main() {
	fmt.Printf("Reading input file\n")
	lines := models.ReadFile(conf.QueryList)
	fmt.Printf("Reading input file [lines=%d]\n", len(lines))

	fmt.Printf("Escaping URLs\n")
	urls := models.ParseLines(lines)
	fmt.Printf("Escaping URLs [lines=%d]\n", len(urls))

	//fmt.Printf("Querying URLs\n")
	//models.QueryURLs(urls)
	//fmt.Printf("Querying URLs\n")

	fmt.Printf("Querying URLs\n")
	models.QueryURLsPaginated(urls)
	fmt.Printf("Querying URLs\n")
}
