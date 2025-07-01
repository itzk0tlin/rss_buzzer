package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello world!\n")

	// wd, _ := os.Getwd()
	// path := string(wd) + string(os.PathSeparator) + "data"
	// os.Mkdir(path, 0755)
	// _, _ = core.FetchRSSFeedFromUrl("https://cubiq.ru/news/feed/", "cubiq")
	// _, _ = core.FetchRSSFeedFromUrl("https://www.opennet.ru/opennews/opennews_all_utf.rss", "opennet")
	// if err != nil {
	// 	panic(err)
	// }
	// rss := core.ParseRSS(feed)
	// fmt.Printf("Fetched rss: %s-%s", rss.Channel.Title, rss.Channel.Link)

	// pairs := core.RSSRepo.GetAllPairs()
	// for _, pair := range pairs {
	// 	fmt.Printf("%s-%s\n", pair.Key, pair.Value)
	// }
	// key, _ := core.RSSRepo.GetKey([]byte("https://cubiq.ru/news/feed/"))
	// fmt.Printf("%s\n", string(key))

	// rss_xml, _ := core.GetRSSFeedFromDB("opennet")
	// rss := core.ParseRSS(rss_xml)
	// for _, item := range rss.Channel.Items {
	// 	fmt.Printf("%s\n------------\n%s\n%s\n\n", item.Title, item.Description, item.PubDate)
	// }
}
