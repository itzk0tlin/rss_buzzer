package core

import (
	"encoding/xml"
	"io"
	"net/http"
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
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
}

func ParseRSS(data []byte) RSS {
	var rss RSS
	xml.Unmarshal(data, &rss)
	return rss
}

func pushRSSLinkInDB(rssTitle []byte, link []byte) error {
	err := RSSRepo.InsertValue(rssTitle, link)
	return err
}

func FetchRSSFeedFromUrl(url string, title string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pushRSSLinkInDB([]byte(title), []byte(url))
	return bytes, nil
}

func GetRSSFeedFromDB(rssTitle string) ([]byte, error) {
	url, err := RSSRepo.GetValue([]byte(rssTitle))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(string(url))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
