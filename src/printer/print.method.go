package printer

import (
	"context"

	"github.com/ciricc/go-grpc-service-template/pkg/genproto/Printer_V1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ser *Service) Print(ctx context.Context, req *Printer_V1.PrintRequest) (
	*Printer_V1.PrintResponse, error,
) {
	if err := validateMessage(req.Message); err != nil {
		return nil, status.Error(codes.Internal, "Invalid message: "+err.Error())
	}
	ser.logger.Println("Printing message: ", req.Message)
	return &Printer_V1.PrintResponse{}, nil
}
