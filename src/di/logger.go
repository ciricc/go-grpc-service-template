package di

import (
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func NewLogger(i *do.Injector) (*logrus.Logger, error) {
	return logrus.New(), nil
}
