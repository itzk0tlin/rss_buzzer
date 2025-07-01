package main

import (
	"fmt"

	"github.com/k0tlin/rss-buzzer/core"
)

func main() {
	fmt.Printf("Hello world!\n")

	// wd, _ := os.Getwd()
	// path := string(wd) + string(os.PathSeparator) + "data"
	// os.Mkdir(path, 0755)
	// feed, err := core.FetchRSSFeed("https://www.opennet.ru/opennews/opennews_all_utf.rss")
	// if err != nil {
	// 	panic(err)
	// }
	// rss := core.ParseRSS(feed)
	// fmt.Printf("Fetched rss: %s", rss.Channel.Title)
	repo := core.RSSRepository{Path: "data"}
	pairs := repo.GetAllPairs()
	for i, pair := range pairs {
		fmt.Printf("%v: %s-%s\n", i, pair.Key, pair.Value)
	}
}
