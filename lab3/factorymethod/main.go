package main

import "fmt"

// startWave принимает ЛЮБОЕ здание (фабрику) и заставляет его произвести бойца.
// Эта функция ничего не знает о конкретных классах, только об интерфейсах Building и Unit.
func startWave(spawner Building) {
	// Делегируем создание объекта конкретной фабрике
	warrior := spawner.SpawnUnit()

	fmt.Printf("В бой вступает: %s\n", warrior.GetName())
	warrior.Action()
	fmt.Println("---------------------------")
}

func main() {
	fmt.Println("=== Подготовка к битве ===")

	// Строим наши фабрики (здания)
	elfFort := &Fort{}
	humanBarracks := &Barracks{}

	// Запускаем производство бойцов.
	// Передаем разные здания в одну и ту же вотфункцию:
	startWave(elfFort)
	startWave(elfFort)
	startWave(humanBarracks)
}
