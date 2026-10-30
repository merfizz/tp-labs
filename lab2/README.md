# Лабораторная работа №2  «Основы ООП»
---

## Описание
В рамках лабораторной работы спроектирована и реализована базовая иерархия игровых объектов для пошаговой стратегии. 

Проект выполнен на языке **Go**. Из-за архитектурных особенностей языка классическое наследование заменено **композицией (встраиванием структур)**, а полиморфизм реализован через **интерфейсы и неявную (утиную) типизацию (duck typing)**.

---

## Архитектура проекта

### Иерархия структур
* `GameObject` — базовый объект с координатами и метаданными.
  * `Unit` — живой объект со здоровьем и возможностью получать урон.
    * `Archer` — лучник (атакует и перемещается).
  * `Building` — постройка со статусом завершения строительства.
    * `Fort` — оборонительное сооружение (атакует).
    * `MobileHouse` — дом на колёсах (перемещается).

### Интерфейсы поведения
Возможности объектов вынесены в отдельные контракты, не привязанные к конкретным веткам иерархии:
* `Attacker` — способность наносить урон.
* `Moveable` — способность изменять координаты на карте.

---

## Принципы ООП в Go

* **Композиция вместо наследования**  
  Повторное использование полей и методов реализовано через анонимное встраивание: `Unit` встраивает `GameObject`, а `Archer` встраивает `Unit`.

* **Инкапсуляция**  
  Внутреннее состояние структур защищено от прямой невалидной мутации извне и управляется исключительно через методы (`ReceiveDamage`, `Move`, геттеры).

* **Полиморфизм (Duck Typing)**  
  Объектам не требуется явно объявлять принадлежность к интерфейсу через ключевые слова вроде `implements`. Если структура реализует методы интерфейса, компилятор Go автоматически считает её удовлетворяющей контракту.

* **Абстрагирование**  
  Поведение (`Attack`, `Move`) отделено от конкретных сущностей и вынесено в независимые интерфейсы (`Attacker`, `Moveable`).
  
**Ход работы:**

1. В соответствии с заданием были определены следующие классы (структуры) и интерфейсы:
   * **GameObject** — базовый класс (структура) всех игровых объектов;
   ```
   type GameObject struct {
	    id   int
	    name string
	    x    int
	    y    int
    }

    //методы класса
   
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
    
   ```

   * **Unit** — класс (структура) игровых юнитов;
   ```
   type Unit struct {
	    GameObject
	    alive bool
	    hp    float32
    }

   //методы класса
   func (u *Unit) IsAlive() bool {
	    return u.alive
    }

    func (u *Unit) GetHp() float32 {
	    return u.hp
    }

    func (u *Unit) ReceiveDamage(damage float32) {
	    u.hp -= damage
	    if u.hp <= 0 {
		    u.alive = false
		    u.hp = 0
	    }
    }
    func (u *Unit) String() string {
	return fmt.Sprintf("Юнит '%s' HP: %.1f, Жив: %t, Позиция: (%d, %d)",
		u.GetName(), u.GetHp(), u.IsAlive(), u.GetX(), u.GetY())
    }

   ```
   * **Building** — класс (структура) построек;
   ```
   type Building struct {
	    GameObject
	    build bool
    }

   //методы класса

    func (b *Building) IsBuilt() bool {
	    return b.build
    }
   ```
   * **Attacker** — интерфейс объектов, способных атаковать;
   * **Moveable** — интерфейс объектов, способных перемещаться.
   ```
   type Attacker interface {
	    Attack(u *Unit)
    }

    type Moveable interface {
	    Move(x, y int)
    }
   ```
   * **Archer** — класс (структура) лучника;
   ```
   type Archer struct {
	    Unit
    }

    func (a *Archer) Attack(target *Unit) {
	    target.ReceiveDamage(15.0)
    }

    func (a *Archer) Move(x, y int) {
	    a.x = x
	    a.y = y
    }
    func (a *Archer) String() string {
	    return fmt.Sprintf("Лучник '%s' HP: %.1f, Жив: %t, Позиция: (%d, %d)",
		    a.GetName(), a.GetHp(), a.IsAlive(), a.GetX(), a.GetY())
    }
   ```
   * **Fort** — класс (структура) крепости;
   ```
   type Fort struct {
	    Building
    }

    func (f *Fort) Attack(target *Unit) {
	    target.ReceiveDamage(30.0)
    }
    func (f *Fort) String() string {
    	return fmt.Sprintf("Крепость '%s' Построено: %t, Позиция: (%d, %d)",
	    	f.GetName(), f.IsBuilt(), f.GetX(), f.GetY())
    }
   ```
   * **MobileHouse** — класс (структура) дома на колёсах;
   ```
   type MobileHouse struct {
	    Building
    }

    func (m *MobileHouse) Move(x, y int) {
	    if m.IsBuilt() {
		    m.x = x
		    m.y = y
	    }
    }
    func (m *MobileHouse) String() string {
    	return fmt.Sprintf("Дом на колесах '%s' Построено: %t, Позиция: (%d, %d)",
	    	m.GetName(), m.IsBuilt(), m.GetX(), m.GetY())
    }
   ```
2. Проверка кода и создание конкретных экземпляров:
```
func main() {
	// Создаем лучника
	archer := &Archer{
		Unit: Unit{
			GameObject: GameObject{id: 1, name: "Merfi", x: 0, y: 0},
			alive:      true,
			hp:         100.0,
		},
	}
	archer.Move(5, 2)
	fmt.Println(archer.GetX(), archer.GetY())
	//5 2

	// Создаем вражеского юнита
	enemy := &Unit{
		GameObject: GameObject{id: 2, name: "Orc", x: 5, y: 5},
		alive:      true,
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
    //проверка атаки и изменения хп
	fort1.Attack(enemy)
	fmt.Println(enemy.GetHp())
	//5

    //проверка смерти при хп < 0
    fort1.Attack(enemy)
	fmt.Println(enemy.GetHp(), enemy.IsAlive())
	//0 false

    //duck typing
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
```
