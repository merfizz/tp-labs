package main

import "fmt"

type Unit struct {
	GameObject
	hp float32
}

// вычисляемая проверка
func (u *Unit) IsAlive() bool {
	return u.hp > 0
}

func (u *Unit) GetHp() float32 {
	return u.hp
}

func (u *Unit) ReceiveDamage(damage float32) {
	u.hp -= damage
	if u.hp <= 0 {
		u.hp = 0
	}
}

func (u *Unit) String() string {
	return fmt.Sprintf("Юнит '%s' HP: %.1f, Жив: %t, Позиция: (%d, %d)",
		u.GetName(), u.GetHp(), u.IsAlive(), u.GetX(), u.GetY())
}
