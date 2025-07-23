# Описание
Проект на Go, реализующий прокси-сервер для моделей медицинских предсказаний (hba1c, ldll, ferr и др).  
Сервер принимает GET-запросы с параметрами, формирует POST-запросы, и перенаправляет их во внешнее API.


# Запуск
1. Go 1.21+
2. Клон:git clone https://github.com/your-username/go-api-practice.git
3. config.yml:
port: ":8080"
auth_token: "Bearer 0l62<EJi/zJx]a?"
predict_api_url: "https://apiml.labhub.online/api/v1/predict"


# Структура проекта
project/
├── cmd/app/main.go              # Точка входа
├── config/config.go             # Загрузка config.yml
├── config.yml                   # Настройки сервера и токен
├── internal/
│   ├── handler/                 # Обработчики HTTP
│   │   └── predict.go
│   ├── services/                # Логика построения JSON-пакетов
│   │   └── predict.go
│   ├── interfaces/              # Интерфейс PredictionBuilder
│   │   └── service.go
│   └── middleware/              # Middleware (авторизация)
│   │   └── auth.go
├── pkg/logger/logger.go         # Логгер
├── logs/app.log                 # Логи (создаются при запуске)
├── go.mod                       # Зависимости
├── go.sum
└── README.md                    # Документация


# Модели
| Модель | URL-путь запроса       | Статус     |
| ------ | ---------------------- | ---------- |
| hba1c  | `/predict?model=hba1c` | ✅ Активен  |
| ldll   | `/predict?model=ldll`  | ✅ Активен  |
| ferr   | `/predict?model=ferr`  | ✅ Активен  |
| ldl    | `/predict?model=ldl`   | ✅ Активен  |
| tg     | `/predict?model=tg`    | ✅ Активен  |
| hdl    | `/predict?model=hdl`   | ❌ Отключен |


# Тестирование
Тестировал с помощью Postman

