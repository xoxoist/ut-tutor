package external

import (
	"github.com/xoxoist/ut-tutor/external/alpha"
	"github.com/xoxoist/ut-tutor/external/bravo"
	"go.uber.org/dig"
)

type DependenciesHolder struct {
	dig.In
	Alpha *alpha.API
	Bravo *bravo.API
}

func RegisterDependencies(container *dig.Container) error {
	var err error
	err = container.Provide(alpha.NewAPI)
	err = container.Provide(bravo.NewAPI)
	if err != nil {
		return err
	}
	return nil
}
