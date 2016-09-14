package main

import (
	"golang.org/x/net/html"
)

func RunParser(config Conf, channel chan HtmlPage){
	config.WG.Add(1)
	defer config.WG.Done()

	anchors := make(chan Anchor)
	go RunChecker(config, anchors)

	for {
		page, more := <- channel

		if !more {
			break
		}

		ParseHtml(anchors, page)
	}

	close(anchors)
}

func ParseHtml(anchors chan Anchor, page HtmlPage){
	z := html.NewTokenizer(page.Stream)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return

		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			for _, a := range t.Attr {
				if a.Key == "href" {
					anchors <- Anchor{a.Val}
					break
				}
			}
		}
	}
}