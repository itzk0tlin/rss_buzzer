package core

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
}

func ParseRSS(data []byte) RSS {
	var rss RSS
	xml.Unmarshal(data, &rss)
	pushRSSLinkInDB([]byte(rss.Channel.Title), []byte(rss.Channel.Link))
	return rss
}

func pushRSSLinkInDB(rssTitle []byte, link []byte) error {
	repo := RSSRepository{Path: "data"}
	err := repo.InsertValue(rssTitle, link)
	return err
}
