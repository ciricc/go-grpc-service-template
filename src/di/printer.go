package di

import (
	"net"

	"github.com/ciricc/go-grpc-service-template/pkg/genproto/Printer_V1"
	"github.com/ciricc/go-grpc-service-template/src/config"
	"github.com/ciricc/go-grpc-service-template/src/pkg/grpcdi"
	"github.com/ciricc/go-grpc-service-template/src/printer"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewPrinterService(i *do.Injector) (*printer.Service, error) {
	logger := do.MustInvoke[*logrus.Logger](i)
	return printer.New(logger)
}

func NewPrinterServer(i *do.Injector) (*grpcdi.Server, error) {
	service := do.MustInvoke[*printer.Service](i)
	config := do.MustInvoke[*config.Config](i)
	logger := do.MustInvoke[*logrus.Logger](i)

	server := grpc.NewServer()

	addr := "0.0.0.0:" + config.Port
	logger.WithFields(logrus.Fields{
		"addr": addr,
	}).Println("serving server")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	Printer_V1.RegisterPrinterServer(server, service)
	return grpcdi.New(server, listener), nil
}
