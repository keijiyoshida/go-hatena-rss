package rss

import "time"

// Item represents item data of the Hatena RSS feed.
type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Date        time.Time `xml:"date"`
}
