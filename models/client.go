package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"tomerg2/go-main/conf"
	"tomerg2/go-main/dtos"
)

func QueryURLs(urls []string) {
	for num, url := range urls {
		fmt.Printf("Processing query [num=%d] [query=%s]\n", num, url)
		queryUrl(url)
	}
}

func QueryURLsPaginated(urls []string) {
	for num, url := range urls {
		fmt.Printf("Processing query [num=%d] [query=%s]\n", num, url)
		totalCount := queryUrlPaginated(url, 1, conf.DefaultPerPage)
		for i := 2; i < (totalCount / conf.DefaultPerPage); i++ {
			queryUrlPaginated(url, i, conf.DefaultPerPage)
		}
	}
}

func queryUrlPaginated(url string, page int, perPage int) (totalCount int) {
	fmt.Printf("Query URL using pagination [page=%d] [per_page=%d]", page, perPage)
	paginatedUrl := AddPagination(url, page, perPage)
	return queryUrl(paginatedUrl)
}

func AddPagination(url string, page int, perPage int) string {
	parsedUrl, _ := url2.Parse(url)
	q := parsedUrl.Query()

	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(perPage))

	parsedUrl.RawQuery = q.Encode()

	fmt.Println(parsedUrl.String())
	return parsedUrl.String()
}

func queryUrl(url string) (totalCount int) {
	var searchResponse dtos.SearchResponse
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", conf.AuthSecret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal([]byte(string(body)), &searchResponse); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	//fmt.Println("Total Count", searchResponse.TotalCount)
	return searchResponse.TotalCount
}
