package di

import (
	"github.com/ciricc/go-grpc-service-template/src/config"
	"github.com/samber/do"
)

func NewConfig(i *do.Injector) (*config.Config, error) {
	c, err := config.New("config.yml")
	return c, err
}
