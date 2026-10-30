package main

import "fmt"

type MobileHouse struct {
	Building
}

func (m *MobileHouse) Move(x, y int) {
	if m.IsBuilt() {
		m.x = x
		m.y = y
	}
}

func (m *MobileHouse) String() string {
	return fmt.Sprintf("Дом на колесах '%s' Построено: %t, Позиция: (%d, %d)",
		m.GetName(), m.IsBuilt(), m.GetX(), m.GetY())
}
