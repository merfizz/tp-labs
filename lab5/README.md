# Лабораторная работа № 5: Структурные паттерны
---
## Описание

В данной лабораторной работе реализованы **структурные паттерны проектирования** на языке **Python**.
В отличие от порождающих (создание) и поведенческих (коммуникация), структурные паттерны решают задачу: **«Как правильно собрать из классов и объектов большие, гибкие и надежные структуры?»**.

---

## 1. Паттерн Proxy

**Суть:**
Предоставляет суррогатный объект (заместитель), который контролирует доступ к реальному объекту, имея с ним абсолютно одинаковый интерфейс.

**Какую проблему решает:**
Позволяет прозрачно для клиента добавить дополнительную логику (кеширование, проверку прав доступа, ленивую загрузку) перед вызовом реального тяжелого объекта. Строго соблюдает принцип **Open/Closed (O)**: мы добавляем кеш, не изменяя код реального загрузчика.

**Кеширующий Прокси:**

```python
import time
from abc import ABC, abstractmethod

# 1. Общий интерфейс (Subject)
class FileDownloader(ABC):
    @abstractmethod
    def download(self, filename: str) -> str:
        pass

# 2. Реальный объект (RealSubject)
class RealDownloader(FileDownloader):
    def download(self, filename: str) -> str:
        print(f"Подключение к серверу... Скачивание файла '{filename}'...")
        time.sleep(2)
        return f"Данные файла {filename}"

# 3. Прокси (Proxy)
class CachedDownloaderProxy(FileDownloader):
    def __init__(self, real_downloader: FileDownloader):
        self._real_downloader = real_downloader
        self._cache: dict[str, str] = {}

    def download(self, filename: str) -> str:
        # Проверяем: если файл уже есть в кеше, отдаем его мгновенно
        if filename in self._cache:
            print(f" Файл '{filename}' найден в кеше")
            return self._cache[filename]
        
        print(f"Файла '{filename}' нет в кеше. Обращение к реальному загрузчику")
        file_data = self._real_downloader.download(filename)
        
        self._cache[filename] = file_data
        return file_data
        
# Клиентский код
def demo_proxy():
    print(" Демонстрация паттерна Прокси")
    
    real_subject = RealDownloader()
    
    proxy = CachedDownloaderProxy(real_subject)

    print("\n Первое обращение:")
    data1 = proxy.download("report_may.pdf")
    
    print("\n Повторное обращение")
    data2 = proxy.download("report_may.pdf")
    
```

---

## 2. Паттерн Adapter

**Суть:**
Выступает в роли «переходника», позволяя объектам с абсолютно несовместимыми интерфейсами работать вместе.

**Какую проблему решает:**
Спасает при интеграции сторонних библиотек (например, Telegram API) в старый код. Адаптер оборачивает чужой объект и переводит вызовы с нашего формата в формат библиотеки, избавляя от необходимости переписывать ядро приложения.

**Система уведомлений:**

```python
# 1. Target 
class NotificationService:
    def send_notification(self, message: str):
        pass

class SMSNotification(NotificationService):
    def send_notification(self, message: str):
        print(f"SMS: Отправка сообщения: '{message}'")

# 2. Adaptee 
class TelegramBotAPI:
    def push_telegram_alert(self, text: str, chat_id: int):
        print(f"Telegram -> Чат {chat_id}: {text}")

# 3. Adapter 
class TelegramAdapter(NotificationService):
    def __init__(self, telegram_api: TelegramBotAPI, chat_id: int):
        self._telegram_api = telegram_api
        self._chat_id = chat_id

    def send_notification(self, message: str):
        self._telegram_api.push_telegram_alert(message, self._chat_id)

# Клиентский код
def alert_user(notifier: NotificationService, text: str):
    notifier.send_notification(text)

def demo_adapter():
    print("\n Демонстрация паттерна Адаптер")
    
    old_sms = SMSNotification()
    alert_user(old_sms, "Ваш заказ собран!")

    telegram_lib = TelegramBotAPI()
    telegram_adapter = TelegramAdapter(telegram_lib, chat_id=100500)
    
    alert_user(telegram_adapter, "Ваш заказ передан курьеру!")


```

---

## 3. Паттерн Мост (Bridge)

**Суть:**
Разделяет одну сущность на две независимые иерархии: Абстракцию (что мы делаем) и Реализацию (с помощью чего мы это делаем), связывая их через композицию.

**Какую проблему решает:**
Предотвращает комбинаторный «взрыв классов». Если бы мы использовали наследование для отчетов (Краткий/Подробный) и способов вывода (Терминал/Файл), нам пришлось бы создавать классы вроде `ShortTerminalReport`. Мост позволяет смешивать виды отчетов и способы вывода на лету.

**Полный код реализации и тестирования:**

```python
from abc import ABC, abstractmethod

# 1. Implementor (Интерфейс реализации)
class ReportSender(ABC):
    @abstractmethod
    def send(self, data: str):
        pass

# 2. ConcreteImplementor (Конкретные реализации)
class TerminalSender(ReportSender):
    def send(self, data: str):
        print(f"Вывод в терминал:\n{data}\n")

class FileSender(ReportSender):
    def __init__(self, filename: str):
        self.filename = filename

    def send(self, data: str):
        print(f"Данные успешно сохранены в файл '{self.filename}'.")

# 3. Abstraction (Базовая абстракция)
class Report(ABC):
    def __init__(self, sender: ReportSender): #композиция 
        self._sender = sender

    @abstractmethod
    def generate(self):
        pass

# 4. RefinedAbstraction (Конкретные абстракции)
class ShortReport(Report):
    def __init__(self, sender: ReportSender, title: str):
        super().__init__(sender)
        self.title = title

    def generate(self):
        content = f"Краткий отчет: {self.title}"
        self._sender.send(content)

class DetailedReport(Report):
    def __init__(self, sender: ReportSender, title: str, details: list[str]):
        super().__init__(sender)
        self.title = title
        self.details = details

    def generate(self):
        content = f"Подробный отчет: {self.title}\n"
        for item in self.details:
            content += f" * {item}\n"
        content += "Конец отчета"
        self._sender.send(content)

# Клиентский код
def demo_bridge():
    print("\n Демонстрация паттерна Мост")
    
    terminal = TerminalSender()
    pdf_file = FileSender("year_report.pdf")

    report1 = ShortReport(terminal, "Продажи за день: 228")
    report2 = DetailedReport(pdf_file, "Годовой отчет", ["Квартал 1: Рост", "Квартал 2: Падение"])

    report1.generate()
    report2.generate()

```