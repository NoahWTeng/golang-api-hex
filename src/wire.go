//go:build wireinject
// +build wireinject

package app

import (
	"api-golang-hex/src/interface/injector"
	"github.com/google/wire"
)

func CreateNewApp() (*AppContainer, error) {
	panic(wire.Build(
		injector.ConfigProvider,
		injector.MongodbProvider,
		wire.Struct(new(AppContainer), "*"),
	))
}
