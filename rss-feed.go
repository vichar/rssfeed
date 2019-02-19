package rssfeed

import (
	"errors"
	"net/http"
	"net/url"
)

func main() {}

//GetContent is function to retrieve Content of an rss feed
func GetContent(feedURL string) (*http.Response, error) {
	r := validateURL(feedURL)
	if r {
		return http.Get(feedURL)
	}

	return nil, errors.New("Invalid URL")
}

func validateURL(feedURL string) bool {
	_, error := url.ParseRequestURI(feedURL)
	if error != nil {
		return false
	}

	return true
}
