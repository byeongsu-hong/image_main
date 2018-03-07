package main

import (
	"./manager"
	"./webtoon"
	"sync"
)

func main() {
	var info = webtoon.GetToonInfo("183559", "mon")

	var wg = new(sync.WaitGroup)

	wg.Add(2)

	var epChan = make(chan webtoon.Episode)
	var done = make(chan bool)

	Webtoon := webtoon.Downloader{
		Wg:     wg,
		EpChan: epChan,
		Done:   done,
		Info:   info,
	}

	Manager := manager.Manager{
		Wg:     wg,
		EpChan: epChan,
		Done:   done,
	}

	go Webtoon.DownloadOne()
	go Manager.Run()

	wg.Wait()
}
