from proxy.downloader import RealDownloader, CachedDownloaderProxy
from adapter.notification import NotificationService, SMSNotification, TelegramBotAPI, TelegramAdapter
from bridge.reports import TerminalSender, FileSender, ShortReport, DetailedReport

def demo_bridge():
    print("\n Демонстрация паттерна Мост")
    
    terminal = TerminalSender()
    pdf_file = FileSender("year_report.pdf")

    report1 = ShortReport(terminal, "Продажи за день: 228")
    report2 = DetailedReport(pdf_file, "Годовой отчет", ["Квартал 1: Рост", "Квартал 2: Падение"])

    report1.generate()
    report2.generate()


def alert_user(notifier: NotificationService, text: str):
    notifier.send_notification(text)

def demo_adapter():
    print("\n Демонстрация паттерна Адаптер")
    
    old_sms = SMSNotification()
    alert_user(old_sms, "Ваш заказ собран!")

    telegram_lib = TelegramBotAPI()
    telegram_adapter = TelegramAdapter(telegram_lib, chat_id=100500)
    
    alert_user(telegram_adapter, "Ваш заказ передан курьеру!")

def demo_proxy():
    print(" Демонстрация паттерна Прокси")
    
    real_subject = RealDownloader()
    
    proxy = CachedDownloaderProxy(real_subject)

    print("\n Первое обращение:")
    data1 = proxy.download("report_may.pdf")
    
    print("\n Повторное обращение")
    data2 = proxy.download("report_may.pdf")
    
if __name__ == "__main__":
    demo_proxy()
    demo_adapter()
    demo_bridge()

