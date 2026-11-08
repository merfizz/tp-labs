package main

import "fmt"

type Archer struct{}

func (a *Archer) GetName() string {
	return "Эльфийский лучник"
}

func (a *Archer) Action() {
	fmt.Println("Стреляет из лука")
}
