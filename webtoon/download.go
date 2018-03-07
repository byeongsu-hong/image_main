package webtoon

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/PuerkitoBio/goquery"
	"log"
	"sync"
	"time"
)

type Downloader struct {
	Info   ToonInfo
	Wg     *sync.WaitGroup
	EpChan chan<- Episode
	Done   chan<- bool
}

func (d *Downloader) DownloadOne() {
	defer func() {
		d.Done <- true
		d.Wg.Done()
	}()

	var err error
	var doc *goquery.Document

	for currentEp := d.Info.MaxNo; 0 < currentEp; currentEp-- {
		var ep Episode

		if doc, err = goquery.NewDocument(DETAIL.parse(d.Info.TitleId, currentEp)); err != nil {
			log.Println(err)
		} else {
			ep.getEpisodeTitle(doc.Selection)
			ep.getAllImageSrc(doc.Selection)
			d.EpChan <- ep
		}

		time.Sleep(time.Duration(randomdata.Number(1000, 2000)) * time.Millisecond)
	}
}
