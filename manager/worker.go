package manager

import (
	"github.com/valyala/gorpc"
)

type worker struct {
	Node *gorpc.Client
}
