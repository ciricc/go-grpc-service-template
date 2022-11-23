package main

import (
	"github.com/ciricc/go-grpc-service-template/src/di"
	"github.com/ciricc/go-grpc-service-template/src/pkg/grpcdi"
	"github.com/samber/do"
)

func main() {
	i := do.New()

	do.Provide(i, di.NewConfig)
	do.Provide(i, di.NewPrinterServer)
	do.Provide(i, di.NewPrinterService)
	do.Provide(i, di.NewLogger)

	server := do.MustInvoke[*grpcdi.Server](i)
	err := server.Start()
	if err != nil {
		panic(err)
	}

	server.Shutdown()

}
