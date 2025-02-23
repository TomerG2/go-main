package models

import (
	"fmt"
	url2 "net/url"
	"strings"
)

func ParseLines(lines []string) (escapedUrls []string) {
	for _, line := range lines {
		escapedUrls = append(escapedUrls, parseLine(line))
	}
	return escapedUrls
}

func parseLine(url string) (escapedUrl string) {
	parts := strings.Split(url, "=")
	escapedUrl = fmt.Sprintf("%s=%s", parts[0], url2.QueryEscape(parts[1]))
	return escapedUrl
}
