from abc import ABC, abstractmethod

# 1. Интерфейс стратегии (Абстрактный класс)
class PaymentStrategy(ABC):
    @abstractmethod
    def pay(self, amount: float) -> bool:
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
