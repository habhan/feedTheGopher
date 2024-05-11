package feeder

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type Feed struct {
	Items []Item `xml:"channel>item"`
}

func (s Item) String() string {
	return fmt.Sprintf("T: %s.\nD: %s.\n", s.Title, s.Description)
}

func Fetcher() {
	g, err := http.Get("https://iamesports.substack.com/feed")
	if err != nil {
		log.Panicln(err)
	}
	slog.Info(g.Status)
	b, err := io.ReadAll(g.Body)
	if err != nil {
		log.Panicln(err)
	}
	f := Feed{}
	err = xml.Unmarshal(b, &f)
	if err != nil {
		log.Panicln(err)
	}
	for _, v := range f.Items {
		fmt.Println("---")
		fmt.Print(v)
	}
}
