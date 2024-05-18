package feeder

import (
	"encoding/xml"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Represents an individual Item
type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Content     string `xml:",any"`
	Link        string `xml:"link"`
}

// Represents the Feed from a channel
type Feed struct {
	Items   []Item `xml:"channel>item"`
	IAuthor string `xml:"itunes:author"`
	GAuthor string `xml:"googleplay:author"`
}

// Takes an url as a string and returns a Feed
func Fetcher(url string) Feed {
	g, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Body.Close()
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
	return f
}

// Takes the HTML input of content and converts it to MD
// TODO: Make this actual MD instead of slice of strings
func MarkDowns(input *string) (result []string) {

	r := strings.NewReader(*input)
	t := html.NewTokenizer(r)
	for {
		token := t.Next()
		switch token {
		case html.ErrorToken:
			return
		case html.TextToken:
			result = append(result, t.Token().Data)
		case html.StartTagToken:
			continue
		case html.EndTagToken:
			continue
		default:
			continue
		}
	}
}
