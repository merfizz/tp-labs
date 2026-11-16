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