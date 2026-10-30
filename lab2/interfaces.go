package main

//абстрактные классы = интерфейсы
type Attacker interface {
	Attack(u *Unit)
}

type Moveable interface {
	Move(x, y int)
}
