package pkg

import "github.com/k0tlin/rss-buzzer/rss/core"

type Client interface {
	DatabasePrefix() string
	SendRSSItem(core.Item, ...any) error
	RegisterRSSFeed(url string, feedName string) error
}
