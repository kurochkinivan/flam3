package variation

import (
	"fmt"
)

type NamedFunction struct {
	Name     VariationName
	Function F
}

func NewNamedFunction(name VariationName, function F) NamedFunction {
	return NamedFunction{
		Name:     name,
		Function: function,
	}
}

type VariationName string

type F func(x, y float64) (newX, newY float64)

func (v VariationName) IsValid() bool {
	_, ok := registry[v]
	return ok
}

func VariationProvider(name VariationName) (NamedFunction, error) {
	if fn, ok := registry[name]; ok {
		return fn, nil
	}

	return NamedFunction{}, fmt.Errorf("unknown variation %q", name)
}
