//go:build wireinject
// +build wireinject

package wire_test

import (
	"context"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(NewStudent, NewClass, NewSchool)

func initSchool(ctx context.Context) (School, error) {
	wire.Build(SuperSet)
	return School{}, nil
}
