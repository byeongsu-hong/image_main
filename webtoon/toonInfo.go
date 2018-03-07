package webtoon

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Day string

const (
	MONDAY    Day = "mon"
	TUESDAY   Day = "tue"
	WEDNESDAY Day = "wed"
	THURSDAY  Day = "thu"
	FRIDAY    Day = "fri"
	SATURDAY  Day = "sat"
	SUNDAY    Day = "sun"
)

var Days = []Day{
	MONDAY,
	TUESDAY,
	WEDNESDAY,
	THURSDAY,
	FRIDAY,
	SATURDAY,
	SUNDAY,
}

func pretty(a interface{}) string {
	b, _ := json.MarshalIndent(a, "", "    ")
	return string(b)
}

type ToonInfo struct {
	Name    string
	Artist  string
	Day     Day
	TitleId string
	MaxNo   int
}

func (a ToonInfo) Pretty() string {
	return pretty(a)
}

type DayToonInfo []ToonInfo

func (a DayToonInfo) Pretty() string {
	return pretty(a)
}

type AllToonInfo map[string]DayToonInfo

func (a AllToonInfo) Pretty() string {
	return pretty(a)
}

func getTitle(s *goquery.Selection) (title string) {
	s.Find("meta").Each(func(_ int, s *goquery.Selection) {
		if prop, ok := s.Attr("property"); ok && prop == "og:title" {
			if t, ok := s.Attr("content"); ok {
				title = t
			} else {
				log.Fatalln("Unexpected no title error")
			}
		}
	})

	return
}

func getTitleId(s *goquery.Selection) string {
	val, ok := s.Find("div a").Attr("href")
	if !ok {
		log.Fatalln("Unexpected no attribute error")
	}

	u, err := url.Parse(val)
	if err != nil {
		log.Fatalln(err)
	}

	return u.Query()["titleId"][0]
}

func getArtist(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("h2 span").Last().Text())
}

func getMaxNo(s *goquery.Selection) int {
	val, ok := s.Find("#content table tbody tr .title a").Attr("href")
	if !ok {
		log.Fatalln("Unexpected no attribute error")
	}

	u, err := url.Parse(val)
	if err != nil {
		log.Fatalln(err)
	}

	var max int
	max, err = strconv.Atoi(u.Query()["no"][0])
	if err != nil {
		log.Fatalln(err)
	}
	return max
}

func GetToonInfo(titleId string, day Day) ToonInfo {
	var err error
	var doc *goquery.Document

	if doc, err = goquery.NewDocument(LIST.parse(titleId)); err != nil {
		log.Fatalln(err)
	}

	var info = ToonInfo{
		Name:    getTitle(doc.Selection),
		Artist:  getArtist(doc.Selection),
		Day:     day,
		TitleId: titleId,
		MaxNo:   getMaxNo(doc.Selection),
	}

	return info
}

func GetDayToonsInfo(day Day) (list DayToonInfo) {
	var err error
	var doc *goquery.Document

	if doc, err = goquery.NewDocument(MAIN.parse(day)); err != nil {
		log.Fatalln(err)
	}

	doc.Find(".daily_img .img_list li").Each(func(_ int, s *goquery.Selection) {
		list = append(list, GetToonInfo(getTitleId(s), day))
	})

	return
}

func GetAllToonsInfo() (list AllToonInfo) {
	list = make(AllToonInfo)

	for _, day := range Days {
		list[string(day)] = GetDayToonsInfo(day)
	}

	return
}
