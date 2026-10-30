package main

import "fmt"

func main() {
	// Создаем лучника
	archer := &Archer{
		Unit: Unit{
			GameObject: GameObject{id: 1, name: "Merfi", x: 0, y: 0},
			hp:         100.0,
		},
	}

	archer.Move(5, 2)
	fmt.Println(archer.GetX(), archer.GetY())
	//5 2

	// Создаем вражеского юнита
	enemy := &Unit{
		GameObject: GameObject{id: 2, name: "Orc", x: 5, y: 5},
		hp:         50.0,
	}

	// Лучник атакует врага
	archer.Attack(enemy)
	fmt.Println(enemy.GetHp())
	//35
	// Перемещаем дом на колесах
	house := &MobileHouse{
		Building: Building{
			GameObject: GameObject{id: 3, name: "Wagon", x: 10, y: 10},
			build:      true,
		},
	}
	house.Move(15, 20)
	fmt.Println(house.GetX(), house.GetY())
	//15 20
	fort1 := &Fort{
		Building: Building{
			GameObject: GameObject{id: 4, name: "Fort", x: 17, y: 7},
			build:      true,
		},
	}

	fort1.Attack(enemy)
	fmt.Println(enemy.GetHp())
	//5

	fort1.Attack(enemy)
	fmt.Println(enemy.GetHp(), enemy.IsAlive())
	//0 false
	var squad []Moveable = []Moveable{
		archer,
		house,
	}

	for _, obj := range squad {
		obj.Move(50, 50)
	}
	fmt.Println(archer.GetX(), archer.GetY())
	//50 50
	fmt.Println(house.GetX(), house.GetY())
	//50 50
	fmt.Println(archer)
	// Лучник 'Merfi' HP: 100.0, Жив: true, Позиция: (50, 50)
	fmt.Println(house)
	//Дом на колесах 'Wagon' Построено: true, Позиция: (50, 50)
	fmt.Println(fort1)
	//Крепость 'Fort' Построено: true, Позиция: (17, 7)
	fmt.Println(enemy)
	//Юнит 'Orc' HP: 0.0, Жив: false, Позиция: (5, 5)
}
