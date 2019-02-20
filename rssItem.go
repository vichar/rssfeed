package rssfeed

import "time"

//RSSItem Object to represent an RSS Item
type RSSItem struct {
	Title       string
	MediaLink   string
	Description string
	Language    string
	PubDate     *time.Time
}
