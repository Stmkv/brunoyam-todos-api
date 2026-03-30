# Todos API

REST API сервис для управления задачами (tasks), написанный на Go с использованием чистой архитектуры.
---

## Возможности сервиса

* CRUD для задач
* JWT авторизация (Bearer Token)
* Swagger документация
* PostgreSQL + pgx
* Миграции базы данных (golang-migrate)
* Docker + docker-compose
* Чистая архитектура (Clean Architecture)

---

## Архитектура

Проект построен по принципам **Clean Architecture**:

```
transport (HTTP / Gin)
        ↓
usecase (application logic)
        ↓
domain (entities + interfaces)
        ↑
repository (Postgres implementation)
```

### Структура проекта

```
cmd/                → точка входа (main.go)
internal/
  config/           → конфигурация
  domain/           → бизнес-сущности и интерфейсы
  usecase/          → бизнес-логика
  repository/       → работа с БД (Postgres)
  transport/        → HTTP handlers (Gin)
  app/              → сервер + миграции
migrations/         → SQL миграции
docs/               → Swagger документация
```

---

## Запуск через Docker

### 1. Собрать и запустить

```bash
docker-compose up --build
```

---

### 2. API будет доступно

```
http://localhost:8080
```

---

### 3. Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## Переменные окружения

Необходимые переменные окружения можно найти в `example.env`

---

## Миграции

Миграции выполняются автоматически при старте приложения.

Расположены в:

```
migrations/
```
