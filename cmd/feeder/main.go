package main

import (
	//"fmt"

	"fmt"

	f "github.com/habhab/feedTheGopher"
)

func main() {
	feed := f.Fetcher("https://iamesports.substack.com/feed")

	for range feed.Items {
		continue
	}
	content := f.MarkDowns(&feed.Items[0].Content)
	for _, s := range content {
		fmt.Println(s)
	}
}
