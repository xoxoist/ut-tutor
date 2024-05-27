package di

import (
	"github.com/xoxoist/ut-tutor/external"
	"github.com/xoxoist/ut-tutor/service"
	"github.com/xoxoist/ut-tutor/shared"
	"go.uber.org/dig"
)

// Container master containers for all dependencies injected
// this global variable will be accessed from main function
// and will provide needed instances across functionalities
var Container = dig.New()

// Injected this struct represents dependencies injections
// bank whole injected instance will be accessed from this
// structure.
type Injected struct {
	Envs      *shared.Envs                // Direct inject
	Externals external.DependenciesHolder // Indirect inject
	Services  service.DependenciesHolder  // Indirect inject
}

// NewInjected initialize dependencies injection entries
// for all dependencies based what this function params
// needed will be injected again using Injected struct.
func NewInjected(
	envs *shared.Envs,
	externals external.DependenciesHolder,
	services service.DependenciesHolder,
) *Injected {
	return &Injected{
		Envs:      envs,
		Externals: externals,
		Services:  services,
	}
}

// init default initialization function from golang
func init() {
	var err error
	// Injecting needed dependencies across functionalities
	err = Container.Provide(shared.NewEnvs)
	err = external.RegisterDependencies(Container)
	err = service.RegisterDependencies(Container)

	// Wrapping up all injected dependencies
	err = Container.Provide(NewInjected)
	if err != nil {
		panic(err)
	}
}
