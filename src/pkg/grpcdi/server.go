package grpcdi

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func New(server *grpc.Server, listener net.Listener) *Server {
	return &Server{
		server:   server,
		listener: listener,
	}
}

func (v *Server) Start() error {
	return v.server.Serve(v.listener)
}

func (v *Server) Shutdown() error {
	v.server.Stop()
	return nil
}
