package service

import (
	"github.com/xoxoist/ut-tutor/service/alpha"
	"github.com/xoxoist/ut-tutor/service/bravo"
	"go.uber.org/dig"
)

type DependenciesHolder struct {
	dig.In
	Alpha alpha.Service
	Bravo bravo.Service
}

func RegisterDependencies(container *dig.Container) error {
	var err error
	err = container.Provide(alpha.NewService)
	err = container.Provide(bravo.NewService)
	if err != nil {
		return err
	}
	return nil
}
