# WebAppBridge 🚀

WebAppBridge — это инструмент для разработчиков Telegram Mini Apps, который упрощает процесс локальной разработки. Он автоматически пробрасывает локальный фронтенд через туннель и обновляет URL у Telegram-бота.

![License](https://img.shields.io/badge/license-MIT-green) 
![Go](https://img.shields.io/badge/Go-1.18%2B-blue)

## ✨ Возможности

- **🚀 Туннелирование** — запускает туннель к локальному серверу через localtunnel.
- **🔗 Публичный URL** — автоматически получает публичный URL для доступа к локальному серверу.
- **🤖 Интеграция с Bot API** — обновляет URL WebApp в настройках Telegram-бота.
- **⚡ Простота использования** — больше не нужно вручную настраивать или менять ссылки.

## 🔧 Установка

Способ 1: Скачивание бинарного файла

Скачайте файл web-app-bridge и запустите файл:
```bash
# Windows
.\web-app-bridge.exe --token YOUR_BOT_TOKEN --port 3000

# macOS (Intel)
mv web-app-bridge-darwin-amd64 web-app-bridge
chmod +x web-app-bridge
./web-app-bridge --token YOUR_BOT_TOKEN --port 3000

# macOS (Apple Silicon/M1/M2)
mv web-app-bridge-darwin-arm64 web-app-bridge
chmod +x web-app-bridge
./web-app-bridge --token YOUR_BOT_TOKEN --port 3000

# Linux
chmod +x web-app-bridge
./web-app-bridge --token YOUR_BOT_TOKEN --port 3000
```

<!-- Скачайте последнюю версию для вашей ОС:
```bash
# Для Linux
wget https://github.com/purkhanov/web-app-bridge/releases/latest/download/web-app-bridge-linux-amd64 -O web-app-bridge
chmod +x web-app-bridge

# Для macOS
wget https://github.com/purkhanov/web-app-bridge/releases/latest/download/web-app-bridge-darwin-amd64 -O web-app-bridge
chmod +x web-app-bridge

# Для Windows
curl -LO https://github.com/purkhanov/web-app-bridge/releases/latest/download/web-app-bridge-windows-amd64.exe
``` -->

Способ 2: Сборка из исходников
```bash
git clone https://github.com/purkhanov/web-app-bridge.git
cd web-app-bridge
go build -o web-app-bridge
```

## 🚀 Использование

Запустите WebAppBridge с токеном бота и портом фронтенда:
```bash
./web-app-bridge --token <BOT_TOKEN> --port <PORT>
```

Пример:
```bash
./web-app-bridge --token 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 --port 3000
```

## 🔄 Что происходит?
1. 🔄 Запуск туннеля — создается безопасное соединение с вашим локальным сервером
2. 🔗 Генерация URL — вы получаете уникальный публичный HTTPS URL (например: https://your-app.loca.lt)
3. 🔐 Получение пароля — генерируется уникальный пароль для доступа к туннелю
4. 🤖 Обновление настроек бота — URL автоматически устанавливается в настройках WebApp вашего Telegram-бота

## 🔐 Работа с паролем

При первом посещении сгенерированного URL:

- **Введите пароль** — который отображается в терминале после запуска WebAppBridge
- **Сохраните доступ** — пароль потребуется только один раз для данного сеанса
- **Безопасность** — пароль защищает ваш локальный сервер от несанкционированного доступа

## 🎯 Для кого
* 👨‍💻 **Фронтенд-разработчики** — создающие Telegram Mini Apps.
* 👥 **Команды** — где важна простая настройка окружения.
* ⚡ **Хакатоны** — быстрая настройка для прототипов и демо
* 🔧 **DevOps** — упрощение процесса разработки и тестирования

## 🤝 Совместимость
* ✅ Telegram Bot API
* ✅ Любой локальный сервер (Vite, React, Vue, Next.js, Express)
* ✅ Windows, macOS, Linux
* ✅ HTTPS с самоподписанными сертификатами

## 📝 Примечания
* Убедитесь, что ваш локальный сервер запущен на указанном порту
* Бот должен иметь права на изменение настроек WebApp
* Пароль требуется только при первом посещении URL

## 📜 Лицензия
MIT License. Подробнее см. в файле [LICENSE](https://github.com/purkhanov/web-app-bridge/blob/main/LICENSE).

## 🐛 Сообщения об ошибках
Нашли баг? [Создайте issue](https://github.com/purkhanov/web-app-bridge/issues) на GitHub.