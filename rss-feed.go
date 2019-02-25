package rssfeed

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mmcdole/gofeed"
)

// ParseRSSData Parsing RSS Data from String
func ParseRSSData(xmlData string) (RSSChannel, error) {
	fp := gofeed.NewParser()
	feed, error := fp.ParseString(xmlData)
	var channel RSSChannel
	if error == nil {
		channel.Title = feed.Title
		channel.Link = feed.Link
		channel.Description = feed.Description
		channel.Language = feed.Language
		channel.LastBuildDate = feed.PublishedParsed
		channel.ImageURL = feed.Image.URL

		items := make([]RSSItem, len(feed.Items))

		for _, item := range feed.Items {
			var rssItem RSSItem
			rssItem.Title = item.Title
			rssItem.Description = item.Description
			rssItem.MediaLink = item.Enclosures[0].URL
			rssItem.PubDate = item.PublishedParsed
			rssItem.Language = channel.Language
			items = append(items, rssItem)
		}

		channel.Items = items
		return channel, nil
	}
	return channel, errors.New("Invalid RSS Feed")
}

//ParseHTTPResponse Parse RSS Content from HTTP Response Body
func ParseHTTPResponse(response *http.Response) HTTPResponse {
	var responseObject HTTPResponse
	responseObject.HTTPStatus = response.Status
	responseObject.HTTPStatusCode = response.StatusCode
	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		responseObject.Body = bodyString

		return responseObject
	}
	responseObject.Body = ""
	return responseObject
}

//HTTPGet is a function to retrieve Content of an rss feed
func HTTPGet(feedURL string) (*http.Response, error) {
	r := validateURL(feedURL)
	if r {
		response, error := http.Get(feedURL)
		return response, error
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
