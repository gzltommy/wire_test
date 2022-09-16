// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire_set

import (
	"context"
	"github.com/google/wire"
)

// Injectors from wire.go:

func initSchool(ctx context.Context) (School, error) {
	student := NewStudent()
	class := NewClass(student)
	school, err := NewSchool(ctx, class)
	if err != nil {
		return School{}, err
	}
	return school, nil
}

// wire.go:

var SuperSet = wire.NewSet(NewStudent, NewClass, NewSchool)