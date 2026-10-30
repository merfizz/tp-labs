package main

// класс = стркутура
// поля приватные
type GameObject struct {
	id   int
	name string
	x    int
	y    int
}

//геттеры
// "g" алтернатива self
func (g *GameObject) GetId() int {
	return g.id
}

func (g *GameObject) GetName() string {
	return g.name
}

func (g *GameObject) GetX() int {
	return g.x
}

func (g *GameObject) GetY() int {
	return g.y
}
