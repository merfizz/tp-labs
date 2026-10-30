package main

import "fmt"

type Fort struct {
	Building
}

func (f *Fort) Attack(target *Unit) {
	target.ReceiveDamage(30.0)
}

func (f *Fort) String() string {
	return fmt.Sprintf("Крепость '%s' Построено: %t, Позиция: (%d, %d)",
		f.GetName(), f.IsBuilt(), f.GetX(), f.GetY())
}
