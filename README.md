# TelePort 🚀

TelePort — это инструмент для разработчиков Telegram Mini Apps, который упрощает процесс локальной разработки. Он автоматически пробрасывает локальный фронтенд через туннель и обновляет URL у Telegram-бота.

![License](https://img.shields.io/badge/license-MIT-green) 
![Go](https://img.shields.io/badge/Go-1.18%2B-blue)

## ✨ Возможности

- **🚀 Туннелирование** — запускает туннель к локальному серверу (через `localtunnel` или `cloudflared`).
- **🔗 Публичный URL** — автоматически получает публичный URL для доступа к локальному серверу.
- **🤖 Интеграция с Bot API** — обновляет URL WebApp в настройках Telegram-бота.
- **⚡ Простота использования** — больше не нужно вручную настраивать `ngrok` или менять ссылки.

## 🔧 Установка

1. Склонируйте репозиторий:
```bash
git clone https://github.com/username/teleport.git
cd teleport