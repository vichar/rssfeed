package rssfeed

import (
	"time"
)

//RSSChannel Object to represent an RSS Feed
type RSSChannel struct {
	Title         string
	Link          string
	Description   string
	Language      string
	ImageURL      string
	LastBuildDate *time.Time
	Items         []RSSItem
}
