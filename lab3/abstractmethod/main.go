package main

import "fmt"

// --- 1. Интерфейсы продуктов (Кого мы создаем) ---

type Warrior interface {
	Attack()
}

type Mage interface {
	CastSpell()
}

// --- 2. Конкретные продукты (Сами бойцы) ---

// Эльфы
type ElfWarrior struct{}

func (e *ElfWarrior) Attack() {
	fmt.Println("- Эльф-воин бьет длинным мечом")
}

type ElfMage struct{}

func (e *ElfMage) CastSpell() {
	fmt.Println("- Эльф-маг стреляет огненным шаром")
}

// Орки
type OrcWarrior struct{}

func (o *OrcWarrior) Attack() {
	fmt.Println("- Орк-воин бьет тяжелой дубиной")
}

type OrcMage struct{}

func (o *OrcMage) CastSpell() {
	fmt.Println("- Орк-шаман кастует цепную молнию")
}

// --- 3. Интерфейс фабрики (Правило для всех баз) ---
// Любая база должна уметь создавать и воина, и мага

type BaseFactory interface {
	CreateWarrior() Warrior
	CreateMage() Mage
}

// --- 4. Конкретные фабрики (Сами базы) ---

// База эльфов
type ElfBase struct{}

func (b *ElfBase) CreateWarrior() Warrior {
	return &ElfWarrior{}
}
func (b *ElfBase) CreateMage() Mage {
	return &ElfMage{}
}

// База орков
type OrcBase struct{}

func (b *OrcBase) CreateWarrior() Warrior {
	return &OrcWarrior{}
}
func (b *OrcBase) CreateMage() Mage {
	return &OrcMage{}
}

// --- 5. Клиентский код (Тестируем всё вместе) ---

// Функция, которая собирает отряд. Ей неважно, чья это база.
func buildSquad(factory BaseFactory) {
	warrior := factory.CreateWarrior()
	mage := factory.CreateMage()

	warrior.Attack()
	mage.CastSpell()
}

func main() {
	fmt.Println("Спавн отряда эльфов:")
	elfBase := &ElfBase{}
	buildSquad(elfBase)

	fmt.Println("\nСпавн отряда орков:")
	orcBase := &OrcBase{}
	buildSquad(orcBase)
}
