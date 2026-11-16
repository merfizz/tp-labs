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