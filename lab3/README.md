# Лабораторная работа №3: Порождающие паттерны проектирования

## Введение
Цель данной лабораторной работы — изучение и практическое применение порождающих паттернов проектирования . Порождающие паттерны решают задачи безопасного и гибкого создания объектов, скрывая логику их инициализации от клиентского кода. 

Особое внимание в работе уделено реализации паттернов с учетом специфики языка Go: использованию "утиной типизации" (duck typing), композиции вместо наследования и инструментов конкурентного программирования (например, `sync.Once`). Это позволяет соблюдать принципы SOLID (в частности, Open/Closed Principle), делая архитектуру расширяемой без модификации существующего кода.

В рамках работы реализованы 4 паттерна:
1. Singleton
2. Factory Method 
3. Abstract Factory 
4. Builder 

---

## 1. Singleton 

### Теоретическая справка
**Суть паттерна:** Гарантирует, что у структуры есть только один экземпляр на всю программу, и предоставляет к нему глобальную точку доступа. 

**Зачем нужен:** Оптимизация ресурсов. Идеально подходит для пулов подключений к базе данных, логгеров или конфигураторов, где многократное создание объектов приведет к утечкам памяти или конфликтам.

**Особенности в Go:** В Go паттерн реализуется потокобезопасно с помощью пакета `sync` и примитива `sync.Once`. Это гарантирует, что даже если тысяча горутин одновременно обратятся к функции создания, объект инициализируется ровно один раз.

### Реализация
```go
package main

import (
	"fmt"
	"sync"
)

// Database — структура, экземпляр которой должен быть один
type Database struct {
	Connection string
}

var (
	instance *Database
	once     sync.Once
)

// GetInstance возвращает единственный экземпляр БД
func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Инициализация подключения к БД... (выполнится лишь 1 раз)")
		instance = &Database{Connection: "Server=localhost;Database=AdventureWorksLT2022;User Id=sa;Password=secret;"}
	})
	return instance
}

func main() {
	db1 := GetInstance()
	db2 := GetInstance()

	if db1 == db2 {
		fmt.Println("Успех: Обе переменные указывают на один и тот же адрес в памяти!")
	}
}
```
## 2. Factory Method 

### Теоретическая справка
**Суть паттерна:** Определяет общий интерфейс для создания объекта, но позволяет подклассам (конкретным фабрикам) самим решать, объект какого типа создавать

**Зачем нужен:** Отвязывает бизнес-логику (клиентский код) от конкретных классов создаваемых продуктов. Позволяет добавлять новые типы продуктов, не меняя существующий код (соблюдение Open/Closed Principle). Клиентский код работает только с абстракциями.

### Реализация
```go
package main

import "fmt"

// 1. Интерфейс Продукта
type Unit interface {
	GetName() string
	Action()
}

// 2. Конкретные Продукты
type Archer struct{}
func (a *Archer) GetName() string { return "Лучник" }
func (a *Archer) Action()         { fmt.Println("Стреляет из лука") }

type Swordsman struct{}
func (s *Swordsman) GetName() string { return "Мечник" }
func (s *Swordsman) Action()         { fmt.Println("Бьет мечом") }

// 3. Интерфейс Фабрики
type Building interface {
	SpawnUnit() Unit
}

// 4. Конкретные Фабрики
type Fort struct{}
func (f *Fort) SpawnUnit() Unit { return &Archer{} }

type Barracks struct{}
func (b *Barracks) SpawnUnit() Unit { return &Swordsman{} }

// Клиентский код (Бизнес-логика)
func startWave(spawner Building) {
	warrior := spawner.SpawnUnit() // Делегирование создания
	fmt.Printf("В бой вступает: %s\n", warrior.GetName())
	warrior.Action()
}

func main() {
	fort := &Fort{}
	barracks := &Barracks{}

	startWave(fort)
	startWave(barracks)
}
```
## 3. Abstract Method 

### Теоретическая справка
**Суть паттерна:** Предоставляет интерфейс для создания целых семейств взаимосвязанных объектов, не привязываясь к их конкретным классам

**В чем отличие от Фабричного метода:** Фабричный метод создает ровно один продукт. Абстрактная фабрика создает комплект продуктов

**Зачем нужен:** Гарантирует, что созданные объекты будут стилистически или логически совместимы друг с другом

### Реализация
```go
package main

import "fmt"

// 1. Интерфейсы продуктов
type Warrior interface { Attack() }
type Mage interface { CastSpell() }

// 2. Конкретные продукты (Семейство Эльфов)
type ElfWarrior struct{}
func (e *ElfWarrior) Attack() { fmt.Println("Эльф бьет мечом") }

type ElfMage struct{}
func (e *ElfMage) CastSpell() { fmt.Println("Эльф кастует фаербол") }

// 2. Конкретные продукты (Семейство Орков)
type OrcWarrior struct{}
func (o *OrcWarrior) Attack() { fmt.Println("Орк бьет дубиной") }

type OrcMage struct{}
func (o *OrcMage) CastSpell() { fmt.Println("Орк кастует молнию") }

// 3. Абстрактная фабрика (Правило для баз: делать и воина, и мага)
type BaseFactory interface {
	CreateWarrior() Warrior
	CreateMage() Mage
}

// 4. Конкретные фабрики
type ElfBase struct{}
func (b *ElfBase) CreateWarrior() Warrior { return &ElfWarrior{} }
func (b *ElfBase) CreateMage() Mage       { return &ElfMage{} }

type OrcBase struct{}
func (b *OrcBase) CreateWarrior() Warrior { return &OrcWarrior{} }
func (b *OrcBase) CreateMage() Mage       { return &OrcMage{} }

// Клиентский код
func buildSquad(factory BaseFactory) {
	warrior := factory.CreateWarrior()
	mage := factory.CreateMage()
	warrior.Attack()
	mage.CastSpell()
}

func main() {
	buildSquad(&ElfBase{})
	buildSquad(&OrcBase{})
}
```
## 4. Builder

### Теоретическая справка
**Суть паттерна:** Позволяет создавать сложные объекты пошагово. Отделяет конструирование объекта от его представления.

**Зачем нужен:** Избавляет от огромного количества параметров при создании. Позволяет использовать один и тот же процесс конструирования для получения разных результатов.

**Особая роль Директора:** Строитель умеет выполнять только отдельные шаги. Именно Директор скрывает в себе правильный алгоритм (порядок вызовов)

### Реализация
```go
package main

import "fmt"

// 1. ПРОДУКТ
type House struct {
	Foundation string 
	Walls      string 
	Roof       string 
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
type Foreman struct {
	builder Builder
}

func (f *Foreman) SetBuilder(b Builder) {
	f.builder = b
}

func (f *Foreman) Construct() {
	fmt.Println("Директор командует: 'Начать стройку по плану!'")
	f.builder.BuildFoundation()
	f.builder.BuildWalls()
	f.builder.BuildRoof()
}


// Клиентский код
func main() {
	// Нанимаем прораба
	foreman := &Foreman{}

	fmt.Println("СТРОИМ ДЕРЕВЯННЫЙ ДОМ")
	woodenTeam := &WoodenBuilder{}
	foreman.SetBuilder(woodenTeam) 
	foreman.Construct()            
	
	woodenHouse := woodenTeam.GetHouse()
	woodenHouse.Show()

	fmt.Println("\nСТРОИМ КАМЕННЫЙ ДОМ")
	stoneTeam := &StoneBuilder{}
	foreman.SetBuilder(stoneTeam) 
	foreman.Construct()
	
	stoneHouse := stoneTeam.GetHouse()
	stoneHouse.Show()
}
```