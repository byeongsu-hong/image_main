package manager

import (
	"github.com/valyala/gorpc"
	"log"
)

var nodeAddr = []string{
	"localhost:1236",
}

type Nodes []*gorpc.Client

func (nds Nodes) Find(key int) *gorpc.Client {
	for k, node := range nds {
		if k == key {
			return node
		}
	}

	log.Fatalln("Unexpected no node in array error")
	return nil
}

func (nds *Nodes) Stop() {
	for _, node := range *nds {
		node.Stop()
	}
}

func (nds *Nodes) Start(title string) {
	for _, node := range *nds {
		node.Start()

		_, err := node.Call("epTitle/" + title)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func listenNodes() (nodes Nodes) {
	for _, node := range nodeAddr {
		nodes = append(nodes, &gorpc.Client{
			Addr: node,
		})
	}
	return
}
