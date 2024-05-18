package feeder

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Content     string `xml:",any"`
	Link        string `xml:"link"`
}

type Feed struct {
	Items   []Item `xml:"channel>item"`
	IAuthor string `xml:"itunes:author"`
	GAuthor string `xml:"googleplay:author"`
}

func (s Item) String() string {
	return fmt.Sprintf("T: %s.\nD: %s.\n", s.Title, s.Description)
}

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

func MarkDowns(input *string) (result []string) {
	//
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
		for _, s := range result {
			fmt.Println(s)
		}
	}
}
