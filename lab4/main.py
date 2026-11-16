# main.py
from strategy.payment import Order, CardPayment, SBPPayment
from chain.validation import DataProvidedValidator, PasswordLengthValidator, UserExistsValidator
from iterator.custom_iterator import Item, Warehouse

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

def demo_strategy():
    print("Демонстрация паттерна Стратегия")
    order1 = Order(order_id=101, amount=1500.0)
    
    order1.set_payment_strategy(CardPayment())
    order1.checkout()

    order2 = Order(order_id=102, amount=3200.0)
    
    order2.set_payment_strategy(SBPPayment())
    order2.checkout()

if __name__ == "__main__":
    demo_strategy()
    demo_chain()
    demo_iterator()