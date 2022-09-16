//go:build wireinject
// +build wireinject

package wire_build

import "github.com/google/wire"

func GetPostService() *PostService {
	panic(wire.Build(
		NewPostService,
		NewPostUseCase,
		NewPostRepo,
	))
}
