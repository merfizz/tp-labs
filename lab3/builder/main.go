package main

import "fmt"

// 1. ПРОДУКТ
type House struct {
	Foundation string // Фундамент
	Walls      string // Стены
	Roof       string // Крыша
}

func (h *House) Show() {
	fmt.Printf("Дом построен! Фундамент: %s | Стены: %s | Крыша: %s\n",
		h.Foundation, h.Walls, h.Roof)
}

// 2. ИНТЕРФЕЙС СТРОИТЕЛЯ
type Builder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	GetHouse() House
}

// 3. КОНКРЕТНЫЙ СТРОИТЕЛЬ №1
type WoodenBuilder struct {
	house House
}

func (w *WoodenBuilder) BuildFoundation() {
	w.house.Foundation = "Деревянные сваи"
	fmt.Println("- Вбиваем деревянные сваи")
}
func (w *WoodenBuilder) BuildWalls() {
	w.house.Walls = "Бревна"
	fmt.Println("- Укладываем сруб из бревен")
}
func (w *WoodenBuilder) BuildRoof() {
	w.house.Roof = "Солома"
	fmt.Println("- Накрываем крышу соломой")
}
func (w *WoodenBuilder) GetHouse() House {
	return w.house
}

// 3. КОНКРЕТНЫЙ СТРОИТЕЛЬ №2
type StoneBuilder struct {
	house House
}

func (s *StoneBuilder) BuildFoundation() {
	s.house.Foundation = "Бетонная плита"
	fmt.Println("- Заливаем надежный бетонный фундамент")
}
func (s *StoneBuilder) BuildWalls() {
	s.house.Walls = "Кирпичи"
	fmt.Println("- Кладем кирпичные стены")
}
func (s *StoneBuilder) BuildRoof() {
	s.house.Roof = "Черепица"
	fmt.Println("- Укладываем черепицу")
}
func (s *StoneBuilder) GetHouse() House {
	return s.house
}

// 4. ДИРЕКТОР
// руководит процессом
type Foreman struct {
	builder Builder
}

// даем бригаду (конкретного строителя)
func (f *Foreman) SetBuilder(b Builder) {
	f.builder = b
}

// Метод управления стройкой.
func (f *Foreman) Construct() {
	fmt.Println("Директор командует: 'Начать стройку по плану!'")
	f.builder.BuildFoundation()
	f.builder.BuildWalls()
	f.builder.BuildRoof()
}

// Клиентский код
func main() {

	foreman := &Foreman{}

	fmt.Println("СТРОИМ ДЕРЕВЯННЫЙ ДОМ")
	woodenTeam := &WoodenBuilder{}
	foreman.SetBuilder(woodenTeam)
	foreman.Construct()

	woodenHouse := woodenTeam.GetHouse()
	woodenHouse.Show()

	fmt.Println("\n СТРОИМ КАМЕННЫЙ ДОМ")
	stoneTeam := &StoneBuilder{}
	foreman.SetBuilder(stoneTeam)
	foreman.Construct()

	stoneHouse := stoneTeam.GetHouse()
	stoneHouse.Show()
}
