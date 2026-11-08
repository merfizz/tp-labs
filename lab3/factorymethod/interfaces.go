package main

// Unit — интерфейс нашего продукта (бойца).
type Unit interface {
	GetName() string
	Action()
}

// Building — интерфейс нашей ФАБРИКИ (создателя).
type Building interface {
	SpawnUnit() Unit
}
