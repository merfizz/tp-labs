package main

import "fmt"

type Fort struct{}

// SpawnUnit — реализация фабричного метода.
func (f *Fort) SpawnUnit() Unit {
	fmt.Println("[Форт]: Обучен новый лучник для защиты стен.")
	return &Archer{} // Фабрика возвращает конкретный тип продукта
}
