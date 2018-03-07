package webtoon

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Episode struct {
	Title  string
	Images []string
}

func (ep *Episode) getEpisodeTitle(s *goquery.Selection) {
	s.Find("meta").Each(func(_ int, s *goquery.Selection) {
		if prop, ok := s.Attr("property"); ok && prop == "og:description" {
			if t, ok := s.Attr("content"); ok {
				ep.Title = t
			} else {
				log.Fatalln("Unexpected no title error")
			}
		}
	})
}

func (ep *Episode) getAllImageSrc(s *goquery.Selection) {
	s.Find("#toonLayer ul li p img").Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("data-original")
		if !ok {
			log.Fatalln("Unexpected no link in img tag error")
		}

		ep.Images = append(ep.Images, val)
	})
}
