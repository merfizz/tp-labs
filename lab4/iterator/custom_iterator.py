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