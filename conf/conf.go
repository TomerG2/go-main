package conf

import "os"

var QueryList string = "/Users/tomer/nanit/repos_backend/go-main/github_queries.txt"
var AuthSecret string = os.Getenv("AuthSecret")
var DefaultPerPage int = 10
