package manager

import (
	"../webtoon"
	"github.com/valyala/gorpc"
	"log"
	"sync"
	"time"
)

type Manager struct {
	Wg     *sync.WaitGroup
	EpChan <-chan webtoon.Episode
	Done   <-chan bool
}

func split(ep webtoon.Episode) {
	var nodes = listenNodes()
	nodes.Start(ep.Title)
	defer nodes.Stop()

	log.Println("Download Start: " + ep.Title)
	var startTime = time.Now()

	var wg = new(sync.WaitGroup)

	var err error
	var cnt = 0
	for _, image := range ep.Images {
		var result *gorpc.AsyncResult
		var node = nodes.Find(cnt)

		result, err = node.CallAsync(image)
		if err != nil {
			log.Fatalln(err)
		}

		wg.Add(1)
		go func() {
			select {
			case <-result.Done:
				wg.Done()
			}
		}()

		cnt += 1
		cnt %= len(nodes)
	}

	wg.Wait()

	log.Println("Download Ended: " + time.Until(startTime).String())
}

func (m *Manager) Run() {
	defer m.Wg.Done()

	for {
		select {
		case ep := <-m.EpChan:
			split(ep)
		case <-m.Done:
			return
		}
	}
}
