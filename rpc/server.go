package rpc

import (
	"net/http"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

type Server struct {
	*service.BaseService

	client rpcclient.Client
	server http.Server
}

// NewServer creates new instance of Server
func NewServer(logger log.Logger) *Server {
	srv := &Server{}
	srv.BaseService = service.NewBaseService(logger, "RPC", srv)
	return srv
}
