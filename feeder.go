package feeder

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	// "golang.org/x/net/html"
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
	// Testingk
	//	d, err := html.Parse(f.Items[0].Content)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//
	//	var links []string
	//	var linkage func(*html.Node)
	//	linkage = func(n *html.Node) {
	//		if n.Type == html.ElementNode && n.Data == "a" {
	//			for _, a := range n.Attr {
	//				if a.Key == "href" {
	//					links = append(links, a.Val)
	//				}
	//			}
	//		}
	//		for c := n.FirstChild; c != nil; c = c.NextSibling {
	//			linkage(c)
	//		}
	//	}
	//	linkage(d)
	//	for _, l := range links {
	//		fmt.Println("Link:", l)
	//	}
	// end testink
}
