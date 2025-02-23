package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		queryUrl(url)
	}
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
	fmt.Println("Total Count", searchResponse.TotalCount)
	return searchResponse.TotalCount
}
