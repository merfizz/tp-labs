package main

import "fmt"

type Archer struct {
	Unit
}

func (a *Archer) Attack(target Damageable) {
	target.ReceiveDamage(15.0)
}

func (a *Archer) Move(x, y int) {
	a.x = x
	a.y = y
}

func (a *Archer) String() string {
	return fmt.Sprintf("Лучник '%s' HP: %.1f, Жив: %t, Позиция: (%d, %d)",
		a.GetName(), a.GetHp(), a.IsAlive(), a.GetX(), a.GetY())

}
