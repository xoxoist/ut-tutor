package main

import (
	"fmt"
	"github.com/xoxoist/ut-tutor/di"
)

func main() {
	err := di.Container.Invoke(func(inj *di.Injected) {
		// get environment variable instance
		envs := inj.Envs
		fmt.Println(envs)
	})
	if err != nil {
		panic(err)
	}
}
