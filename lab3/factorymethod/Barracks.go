package main

import "fmt"

type Barracks struct{}

// SpawnUnit — реализация фабричного метода.
func (b *Barracks) SpawnUnit() Unit {
	fmt.Println("[Казармы]: Выкован меч и обучен новый боец.")
	return &Swordsman{}
}
