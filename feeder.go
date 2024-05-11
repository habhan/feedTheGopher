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
	Items []Item
}

func (s Item) String() string {
	return fmt.Sprintf("Title: %s\nDescrtiption: %s", s.Title, s.Description)
}

func Fetcher() {
	g, err := http.Get("https://iamesports.substack.com/feed")
	if err != nil {
		log.Panicln(err)
	}
	slog.Info(g.Status)
	b, err := io.ReadAll(g.Body)
	i := Item{"", ""}
	if err != nil {
		log.Panicln(err)
	}
	err = xml.Unmarshal(b, &i)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(i)

}
