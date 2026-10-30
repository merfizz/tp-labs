package main

import "fmt"

type Unit struct {
	GameObject
	alive bool
	hp    float32
}

func (u *Unit) IsAlive() bool {
	return u.alive
}

func (u *Unit) GetHp() float32 {
	return u.hp
}

func (u *Unit) ReceiveDamage(damage float32) {
	u.hp -= damage
	if u.hp <= 0 {
		u.alive = false
		u.hp = 0
	}
}

func (u *Unit) String() string {
	return fmt.Sprintf("Юнит '%s' HP: %.1f, Жив: %t, Позиция: (%d, %d)",
		u.GetName(), u.GetHp(), u.IsAlive(), u.GetX(), u.GetY())
}
