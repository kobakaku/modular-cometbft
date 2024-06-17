package rpc

import (
	"net"
	"net/http"
	"strings"

	cmcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"

	"github.com/kobakaku/modular-cometbft/node"
)

type Server struct {
	*service.BaseService

	config *cmcfg.RPCConfig

	client rpcclient.Client
	server http.Server
}

// NewServer creates new instance of Server
func NewServer(node node.Node, config *cmcfg.RPCConfig, logger log.Logger) *Server {
	srv := &Server{
		client: node.GetClient(),
		config: config,
	}
	srv.BaseService = service.NewBaseService(logger, "RPC", srv)
	return srv
}

func (s *Server) OnStart() error {
	return s.startRPC()
}

func (s *Server) OnStop() {}

func (s *Server) startRPC() error {
	parts := strings.SplitN(s.config.ListenAddress, "://", 2)
	proto := parts[0]
	addr := parts[1]

	listener, err := net.Listen(proto, addr)
	if err != nil {
		return nil
	}

	go func() { s.serve(listener) }()

	return err
}

func (s *Server) serve(listener net.Listener) error {
	s.Logger.Info("serving HTTP", "listen address", listener.Addr())
	s.server = http.Server{}
	return s.server.Serve(listener)
}
