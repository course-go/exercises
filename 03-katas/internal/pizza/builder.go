package pizza

import (
	"errors"
	"slices"
)

// Builder is a traditional builder pattern as you know it.
type Builder struct {
	pizza Pizza
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() (pizza *Pizza, err error) {
	if b.pizza.Size == "" {
		return nil, errors.New("size was not selected")
	}

	return &b.pizza, nil
}

func (b *Builder) AddTopping(topping Topping) *Builder {
	if slices.Contains(b.pizza.Toppings, topping) {
		return b
	}

	b.pizza.Toppings = append(b.pizza.Toppings, topping)
	return b
}

func (b *Builder) SetBase(base Base) *Builder {
	b.pizza.Base = base
	return b
}

func (b *Builder) SetSize(size Size) *Builder {
	b.pizza.Size = size
	return b
}
