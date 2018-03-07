package webtoon

import (
	"fmt"
)

type link string

const (
	MAIN   link = "http://comic.naver.com/webtoon/weekdayList.nhn?week=%s"
	LIST   link = "http://comic.naver.com/webtoon/list.nhn?titleId=%s"
	DETAIL link = "http://m.comic.naver.com/webtoon/detail.nhn?titleId=%s&no=%d"
)

func (l link) parse(args ...interface{}) string {
	return fmt.Sprintf(string(l), args...)
}
