package main

//абстрактные классы = интерфейсы
// type Attacker interface {
// 	Attack(u *Unit)
// }
//устарняем зависимость интерфейса от конкретного юнита,
// поэтому добавили интрфейс цели
type Damageable interface {
	ReceiveDamage(damage float32)
}

type Attacker interface {
	Attack(target Damageable)
}
type Moveable interface {
	Move(x, y int)
}
