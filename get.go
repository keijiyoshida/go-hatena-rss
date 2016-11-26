package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://b.hatena.ne.jp/entrylist/%s?mode=rss&sort=%s"

// Get gets and parses the Hatena RSS feed and returns the result.
func Get(category string, sortKey string) ([]Item, error) {
	data, err := get(category, sortKey)
	if err != nil {
		return nil, err
	}

	return parse(data)
}

func get(category string, sortKey string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(url, category, sortKey))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parse(data []byte) ([]Item, error) {
	var rdf RDF

	if err := xml.Unmarshal(data, &rdf); err != nil {
		return nil, err
	}

	return rdf.Items, nil
}
