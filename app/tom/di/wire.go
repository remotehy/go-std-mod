//go:build wireinject
// +build wireinject

package di

import (
	"github.com/remotehy/go-std-mod/app/tom/service"

	"github.com/google/wire"
)

func NewService() service.Service {
	panic(wire.Build(service.Provider))
}
