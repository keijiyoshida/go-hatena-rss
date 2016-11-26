package rss

// RDF represents RDF data of the Hatena RSS feed.
type RDF struct {
	Items []Item `xml:"item"`
}
