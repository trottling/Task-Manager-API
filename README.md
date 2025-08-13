# Task Manager API (Go)

Лёгкое REST API для задач на Go: in-memory хранилище, валидация, логирование, graceful shutdown, Swagger-доки из YAML.

## Фичи

- Создать задачу - POST /tasks
- Получить задачу по ID - GET /tasks/{id}
- Список по статусу + пагинация - GET /tasks?status=...&limit=...&offset=...
- Валидация входных данных
- Асинхронный логгер с middleware на запросы
- Graceful shutdown

## Стек

- Go 1.22+
- Весь проект на встроенных библиотеках

## Структура

```

.
├── cmd/app                 # main.go (запуск сервера)
├── internal
│   ├── api
│   │   ├── dto            # DTO запросов и ответов
│   │   ├── handlers       # Router-структура + хендлеры + встроенный Swagger UI
│   │   ├── mappers        # model -> DTO
│   │   └── validation     # валидация DTO
│   ├── core
│   │   ├── logger         # асинхронный логгер + middleware
│   │   ├── models         # доменные модели (Task, Status)
│   │   └── storage        # in-memory Storage с RWMutex
│   └── utils              # WriteJSON, graceful run
├── Makefile
├── Dockerfile
└── go.mod

```

## Запуск локально

```bash
go run ./cmd/app
# или
make run
````

Сервер слушает `:8080`.

## Эндпоинты

### POST /tasks

Создать задачу:

`Request`

```json
{
  "name": "Buy milk",
  "status": "new"
}
```

`201 response`

```json
{
  "id": 1,
  "name": "Buy milk",
  "status": "new",
  "created_at": "2025-08-13T12:34:56Z"
}
```

### GET /tasks/{id}

`Получить задачу по ID:`

`200 response`

```json
{
  "id": 1,
  "name": "Buy milk",
  "status": "new",
  "created_at": "2025-08-13T12:34:56Z"
}
```

### GET /tasks?status=new\&limit=10\&offset=0

Список задач по статусу:

`200 response`

```json{
  "items": [
    { "id": 1, "name": "Buy milk", "status": "new", "created_at": "..." }
  ],
  "total": 1
}
```

Статусы: `new`, `in_progress`, `done`, можно изменить в `internal/core/models/task.go`, не забудьте обновить валидацию в `internal\api\validation\validator.go`

## Тесты

Прогнать все тесты:

```bash
make test
# под капотом: go test ./... -v
```

## Сборка

```bash
make build
# бинарник: bin/server
```

## Docker

Собрать и запустить:

```bash
make docker-build
make docker-run
# или
docker build -t task-api:latest .
docker run --rm -p 8080:8080 task-api:latest
```

## Graceful shutdown

Сервис ловит SIGINT/SIGTERM, перестает принимать новые соединения и ждёт активные запросы до таймаута (по умолчанию 15s). Логи старт/стоп и ошибки выводятся через асинхронный логгер.

## Примеры curl

```bash
# создать
curl -s -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"name":"Buy milk","status":"new"}' | jq

# получить по id
curl -s http://localhost:8080/tasks/1 | jq

# список по статусу
curl -s 'http://localhost:8080/tasks?status=new&limit=10&offset=0' | jq
```

## Нюансы реализации

* Валидация и нормализация query: limit - дефолт 50, максимум 200, offset - неотрицательный.
* Хранилище - RWMutex, пагинация на срезе после фильтра.
* Логирование:

    * middleware пишет метод, путь, статус и длительность
    * хендлеры логируют валидаторные 4xx как Info, 5xx как Error, бизнес-события - Info

```
