package variation

import (
	"fmt"
)

type NamedFunction struct {
	Name     Name
	Function F
}

func NewNamedFunction(name Name, function F) NamedFunction {
	return NamedFunction{
		Name:     name,
		Function: function,
	}
}

type Name string

type F func(x, y float64) (newX, newY float64)

func (v Name) IsValid() bool {
	_, ok := registry[v]
	return ok
}

func Provider(name Name) (NamedFunction, error) {
	if fn, ok := registry[name]; ok {
		return fn, nil
	}

	return NamedFunction{}, fmt.Errorf("unknown variation %q", name)
}
