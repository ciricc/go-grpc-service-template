package printer

import (
	"github.com/ciricc/go-grpc-service-template/pkg/genproto/Printer_V1"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Printer_V1.UnimplementedPrinterServer
	logger *logrus.Logger
}

func New(logger *logrus.Logger) (*Service, error) {
	return &Service{
		logger: logger,
	}, nil
}
