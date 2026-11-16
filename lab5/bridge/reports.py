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