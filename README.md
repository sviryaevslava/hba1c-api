# API-прокси для ML-моделей (hba1c, ldl, tg, ferr, hdl)

Go-приложение-прокси, принимающее GET-запросы с параметрами и перенаправляющее их на сторонние API медицинских предсказаний. Ответ возвращается в виде JSON-клиенту.

# Структура проекта

GO-API-PRACTICE/
├── cmd/
│   └── app/
│       └── main.go              # Точка входа, настройка сервера
├── config/
│   └── config.go                # Загрузка и парсинг config.yml
├── config.yml                   # Конфигурация сервера и URL-ов
├── internal/
│   ├── handler/
│   │   └── predict.go           # HTTP-хендлер маршрута /predict
│   ├── interfaces/
│   │   └── service.go           # Интерфейс PredictionBuilder
│   ├── middleware/
│   │   └── auth.go              # Middleware для авторизации
│   ├── models/                  # Структуры данных моделей (пусто)
│   └── services/
│       └── predict.go           # Формирование и отправка POST-запроса
├── pkg/
│   └── logger/
│       └── logger.go            # Логгер
├── logs/
│   └── app.log                  # Лог-файл (создается при запуске)
├── .gitignore                   # Игнорируемые файлы Git
├── .dockerignore               # Исключения при сборке Docker
├── docker-compose.yml          # Сборка и запуск через Docker Compose
├── Dockerfile                  # Docker-образ с многоэтапной сборкой
├── go.mod / go.sum             # Модули и зависимости
├── app.exe                     # Бинарник (при локальной сборке)
└── README.md                   # Документация (этот файл)


# Конфигурация
port: ":8080"
auth_token: "0l62<EJi/zJx]a?"
api_urls:
  hba1c: "https://apiml.labhub.online/api/v1/predict/hba1c"
  ldll: "https://apiml.labhub.online/api/v1/predict/ldll"
  ferr: "https://apiml.labhub.online/api/v1/predict/ferr"
  ldl: "https://apiml.labhub.online/api/v1/predict/ldl"
  tg: "https://apiml.labhub.online/api/v1/predict/tg"
  hdl: "https://apiml.labhub.online/api/v1/predict/hdl"

# Запуск

## Локально (без Docker)

1. Установи Go 1.21+
2. Установи зависимости:
go mod tidy
3. Сборка:
go build -o app.exe ./cmd/app
4. Запуск:
./app.exe

## Docker
1. Сборка и запуск
docker-compose up --build -d
2. Остановка
docker-compose down

# Эндпоинты
Все запросы требуют заголовки:
Authorization: Bearer 0l62<EJi/zJx]a?
Accept: application/json; charset=utf-8

# Примеры эндпоинтов:

| Модель | URL                        | Статус     |
| ------ | -------------------------- | ---------- |
| hba1c  | `/predict?model=hba1c&...` | ✅ Активен  |
| ldll   | `/predict?model=ldll&...`  | ✅ Активен  |
| ferr   | `/predict?model=ferr&...`  | ✅ Активен  |
| ldl    | `/predict?model=ldl&...`   | ✅ Активен  |
| tg     | `/predict?model=tg&...`    | ✅ Активен  |
| hdl    | `/predict?model=hdl&...`   | ❌ Отключен |

# Тестирование
Через Postman:
* Метод: `GET`
* URL: `http://localhost:8080/predict?model=hba1c&...`
* Заголовки:

  * `Authorization: Bearer 0l62<EJi/zJx]a?`
  * `Accept: application/json; charset=utf-8`
* Body: отсутствует (используются query-параметры)
