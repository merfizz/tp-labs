# Лабораторная работа № 4: Поведенческие паттерны
---
## Описание
В данной лабораторной работе реализованы поведенческие паттерны проектирования на языке **Python**. 
В отличие от структурных паттернов, которые собирают объекты в большие структуры, поведенческие паттерны решают задачу эффективной и безопасной **коммуникации между объектами**, а также грамотного распределения обязанностей.

---

## 1. Паттерн «Стратегия» (Strategy)

**Суть:**
Паттерн определяет семейство схожих алгоритмов, инкапсулирует каждый из них в собственный класс и делает их взаимозаменяемыми прямо во время выполнения программы.

**Какую проблему решает:**
Избавляет от огромных конструкций `if/else` при выборе поведения (например, выбора способа оплаты). Позволяет клиенту (Контексту) менять алгоритм работы на лету, не меняя при этом собственный код.

**Система оплаты:**
```python
from abc import ABC, abstractmethod

# 1. Общий интерфейс стратегии
class PaymentStrategy(ABC):
    @abstractmethod
    def pay(self, amount: float):
        pass

# 2. Конкретные стратегии 
class SBPPayment(PaymentStrategy):
    def pay(self, amount: float) -> bool:
        print(f"Оплата {amount} ₽ через СБП прошла успешно.")
        return True

class CryptoPayment(PaymentStrategy):
    def pay(self, amount: float) -> bool:
        print(f"Перевод эквивалента {amount} ₽ в USDT завершен.")
        return True

class CardPayment(PaymentStrategy):
    def pay(self, amount: float) -> bool:
        print(f"Списание {amount} ₽ с банковской карты выполнено.")
        return True


# 3. Контекст
class Order:
    def __init__(self, order_id: int, amount: float):
        self.order_id = order_id
        self.amount = amount
        self._payment_strategy = None

    def set_payment_strategy(self, strategy: PaymentStrategy):
        self._payment_strategy = strategy

    def checkout(self):
        if not self._payment_strategy:
            print(f"Ошибка заказа #{self.order_id}: Выберите способ оплаты!")
            return
        
        print(f"\n Оформление заказа #{self.order_id}...")
        self._payment_strategy.pay(self.amount)

# Клиентский код
def demo_strategy():
    print("Демонстрация паттерна Стратегия")
    order1 = Order(order_id=101, amount=1500.0)
    
    order1.set_payment_strategy(CardPayment())
    order1.checkout()

    order2 = Order(order_id=102, amount=3200.0)
    
    order2.set_payment_strategy(SBPPayment())
    order2.checkout()
```

---

## 2. Паттерн «Цепочка обязанностей» (Chain of Responsibility)

**Суть:**
Позволяет передавать запросы последовательно по цепочке обработчиков. Каждый обработчик решает, может ли он обработать запрос сам, должен ли передать его дальше, или обязан прервать цепочку.

**Какую проблему решает:**
Снижает связанность (отправитель не знает, кто именно обработает запрос). Идеально подходит для реализации систем фильтрации и поэтапной валидации. Позволяет добавлять новые проверки не переписывая старый код.

**Валидация регистрации:**
```python
from abc import ABC, abstractmethod

# 1. Интерфейс обработчика
class Validator(ABC):
    @abstractmethod
    def set_next(self, validator: 'Validator') -> 'Validator':
        pass

    @abstractmethod
    def check(self, data: dict) -> bool:
        pass

# 2. Базовый обработчик 
class BaseValidator(Validator):
    def __init__(self):
        self._next_validator = None

    def set_next(self, validator: 'Validator') -> 'Validator':
        self._next_validator = validator
        return validator 

    def check(self, data: dict) -> bool:
        if self._next_validator:
            return self._next_validator.check(data)
        return True 

# 3. Конкретные звенья цепи
class DataProvidedValidator(BaseValidator):
    def check(self, data: dict) -> bool:
        if not data.get("username") or not data.get("password"):
            print("Ошибка: Логин и пароль не могут быть пустыми!")
            return False 
        print("Проверка 1 пройдена: Данные заполнены.")
        return super().check(data) # Передаем следующему

class PasswordLengthValidator(BaseValidator):
    def check(self, data: dict) -> bool:
        if len(data.get("password", "")) < 8:
            print("Ошибка: Пароль слишком короткий (минимум 8 символов)!")
            return False
        print("Проверка 2 пройдена: Длина пароля достаточная.")
        return super().check(data)

class UserExistsValidator(BaseValidator):
    def check(self, data: dict) -> bool:
        db_users = ["admin", "student123"]
        if data.get("username") in db_users:
            print(f"Ошибка: Имя пользователя '{data.get('username')}' уже занято!")
            return False
        print("Проверка 3 пройдена: Логин свободен.")
        return super().check(data)

# Клиентский код
def demo_chain():
    print("\n Демонстрация паттерна Цепочка обязанностей")
    
    registration_chain = DataProvidedValidator()
    registration_chain.set_next(PasswordLengthValidator()).set_next(UserExistsValidator())

    print("\n Пустые данные")
    registration_chain.check({"username": "", "password": "123"})

    print("\n  Короткий пароль")
    registration_chain.check({"username": "new_user", "password": "123"})

    print("\n Занятый логин")
    registration_chain.check({"username": "admin", "password": "superpassword"})

    print("\n Идеальный запрос")
    is_valid = registration_chain.check({"username": "new_user", "password": "superpassword"})
    
    if is_valid:
        print("Валидация успешна!")

```

---

## 3. Паттерн «Итератор» (Iterator)

**Суть (для защиты):**
Предоставляет способ последовательного обхода элементов сложной коллекции, полностью скрывая от клиента её внутреннее устройство (список, дерево, граф или БД).

**Какую проблему решает:**
В Python реализуется через переопределение магических методов `__iter__()` и `__next__()`. Забирает на себя ответственность за сложную логику фильтрации, экономя оперативную память (ленивая загрузка), и позволяет клиенту использовать стандартный цикл `for ... in ...`.

**Фильтрация товаров на складе:**
```python
# 1. Базовая сущность
class Item:
    def __init__(self, name: str, in_stock: bool):
        self.name = name
        self.in_stock = in_stock

    def __str__(self):
        status = "В наличии" if self.in_stock else "Нет на складе"
        return f"Товар: {self.name} [{status}]"


# 2. Конкретный Итератор
class AvailableItemsIterator:
    def __init__(self, items: list[Item]):
        self._items = items
        self._index = 0

    def __iter__(self):
        return self

    def __next__(self) -> Item:
        while self._index < len(self._items):
            current_item = self._items[self._index]
            self._index += 1
            
            if current_item.in_stock:
                return current_item
                
        raise StopIteration()


# 3. Сама коллекция (Агрегат)
class Warehouse:
    def __init__(self):
        self._items: list[Item] = []

    def add_item(self, item: Item):
        self._items.append(item)

    def __iter__(self):
        return AvailableItemsIterator(self._items)
# Клиентский код
def demo_iterator():
    print("\n Демонстрация паттерна Итератор")
    
    warehouse = Warehouse()
    warehouse.add_item(Item("Ноутбук", in_stock=True))
    warehouse.add_item(Item("Мышка", in_stock=False)) 
    warehouse.add_item(Item("Клавиатура", in_stock=True))
    warehouse.add_item(Item("Монитор", in_stock=False))

    print("Каталог доступных товаров для клиента:")
    
    for item in warehouse:
        print(f" - {item.name}")


```