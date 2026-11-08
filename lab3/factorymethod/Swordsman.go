package main

import "fmt"

type Swordsman struct{}

func (s *Swordsman) GetName() string {
	return "Тяжелый мечник"
}

func (s *Swordsman) Action() {
	fmt.Println("Бьет мечом")
}
