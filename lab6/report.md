# Отчёт по лабораторной работе №6

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Собрать стек front + back + db через Docker Compose, проверить взаимодействие сервисов, вставку данных в PostgreSQL и сохранность данных в именованном томе после пересоздания контейнеров.

## Ход работы

Файлы проекта в `work`: `compose.yaml`, `Dockerfile`, `main.go`, `default.conf`, `index.html`, `users.sql`, `.env`.

Запуск и обслуживание:

```bash
sudo docker compose --project-name lab6app up -d --build
sudo docker compose --project-name lab6app ps
wget -qO- http://127.0.0.1:18888
wget -qO- http://127.0.0.1:18888/back
```

Вставка пользователя в БД:

```bash
sudo docker exec lab6app-db-1 psql -U app -d app -c "insert into users (name, email) values ('alex', 'alex@mail.ru');"
```

Проверка изоляции сети (front не должен достучаться до db по имени):

```bash
sudo docker exec lab6app-front-1 sh -lc 'wget -T 2 -qO- http://db:5432'
```

Пересоздание без тома:

```bash
sudo docker compose --project-name lab6app rm -sf
sudo docker compose --project-name lab6app up -d
```

Полный лог сессии: `artifacts/logs/session_lab6.typescript`, выборочные ответы — `front_index.out`, `back_*.json`, `front_to_db.err`.

## Результаты (фрагменты вывода)

Ответ `/back` после вставки третьей записи (файл `artifacts/logs/back_after_insert.json`):

```json
[{"id":1,"name":"john","email":"john@example.com"},{"id":2,"name":"mary","email":"mary@example.com"},{"id":3,"name":"alex","email":"alex@mail.ru"}]
```

Попытка front → db:

```text
wget: bad address 'db:5432'
```

## Замечания

Порт фронта **18888** вместо 8888 из-за занятого порта на хосте. Образы взяты актуальные alpine-варианты nginx и postgres — поведение как в методичке.

## Вывод

Compose поднимает три сервиса, фронт проксирует к backend, backend читает Postgres. Запись `alex` видна в JSON. После `rm -sf` и повторного `up` данные остаются в томе `db-data`. Схема сетей соблюдена: из front до db по DNS не достучаться.
