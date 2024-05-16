package main

import (
	"fmt"

	f "github.com/habhab/feedTheGopher"
)

func main() {
	feed := f.Fetcher("https://iamesports.substack.com/feed")

	for _, i := range feed.Items {
		fmt.Println(i)
	}
}
