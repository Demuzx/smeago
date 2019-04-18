package smeago

import (
	"golang.org/x/net/html"
	"io"
)

type Result struct {
	Links []string
}

func ReadStringSize(rd io.Reader) (*Result, error) {
	r := &Result{}
	links := getLinks(rd)
	lc := links[:0]
	// Only internal links
	for _, l := range links {
		if l[0] == '/' {
			lc = append(lc, decodeURL(l))
		}
	}
	r.Links = lc
	return r, nil
}

func ReadString(rd io.Reader) (*Result, error) {
	r := &Result{}
	links := getLinks(rd)
	lc := links[:0]
	// Only internal links
	for _, l := range links {
		if l[0] == '/' {
			lc = append(lc, decodeURL(l))
		}
	}
	r.Links = lc
	return r, nil
}

func decodeURL(s string) string {
	return html.UnescapeString(s)
}

func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
