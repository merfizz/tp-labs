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