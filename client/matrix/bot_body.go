package matrix

import "github.com/k0tlin/rss-buzzer/rss/core"

type MatrixBot struct {
	homeserver string
	username   string
	password   string
	roomIDs    []int

	databasePrefix string
}

func (bot MatrixBot) DatabasePrefix() string {
	return bot.databasePrefix
}

func (bot MatrixBot) SendRSSItem(rssItem core.Item, roomID int) error {

}

func (bot MatrixBot) RegisterRSSFeed(url string, feedName string) error {
	_, err := core.FetchRSSFeedFromUrl(url, feedName)
	return err
}
