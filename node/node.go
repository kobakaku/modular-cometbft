package node

import (
	"github.com/cometbft/cometbft/libs/log"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

type Node interface {
	Start() error
	GetClient() rpcclient.Client
}

func NewNode(logger log.Logger) (Node, error) {
	// TODO: light or full nodeで切り替えられるようにする
	return newLightNode(logger)
}
