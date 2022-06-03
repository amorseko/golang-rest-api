//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedServer() *SimpleServices {
	wire.Build(SimpleRepository{}, SimpleServices{})
	return nil
}
