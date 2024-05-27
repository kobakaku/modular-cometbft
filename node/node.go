package node

import (
	"github.com/cometbft/cometbft/libs/log"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/kobakaku/modular-cometbft/config"
)

type Node interface {
	Start() error
	GetClient() rpcclient.Client
	Stop() error
	IsRunning() bool
}

func NewNode(conf config.NodeConfig, logger log.Logger) (Node, error) {

	if conf.Light {
		return newLightNode(logger)
	} else {
		return newFullNode(logger)
	}
}
