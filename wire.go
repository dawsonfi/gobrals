//go:build wireinject
// +build wireinject

package main

import (
	"github.com/dawsonfi/gobrals/internal/address"
	"github.com/google/wire"
)

func InitializeAddressHandler(dbPath string) (*address.Handler, error) {
	wire.Build(address.NewHandler, address.NewService, address.NewJsonDatabaseLoader)

	return &address.Handler{}, nil
}
