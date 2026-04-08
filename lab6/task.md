# Docker Compose

**Источник курса:** [06_practice_docker_compose.html](https://yudolevich.github.io/infra-course/06_practice_docker_compose.html)

## Задание

В данном практическом занятии ознакомиться со сборкой и развёртыванием среды из нескольких связанных контейнеров при помощи [docker compose](https://docs.docker.com/compose/).

### Vagrant

```ruby
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/lunar64"
  config.vm.provision "docker"
  config.vm.network "forwarded_port", guest: 8888, host: 8888
end
```

С `docker` provisioner на машину ставится **docker compose** как плагин; либо установка по [инструкции Docker](https://docs.docker.com/compose/install/linux/#install-the-plugin-manually).

### Project

Файл `compose.yaml` с сервисом `front` (`nginx:1.25`, порт 8888:80, `restart: always`). Запуск: `docker compose up -d`. Проверка: `curl -s localhost:8888 | grep title`.

Подкоманды: `docker compose ls`, `docker compose ps`, `docker compose top`.

### Front

Статическая `index.html` (таблица пользователей, запрос к `/back`) и `default.conf` для nginx с `location /back` и `resolver 127.0.0.11`, `proxy_pass` на сервис `back`.

Обновлённый `compose.yaml`: `volumes` для `index.html`, `configs` для `default.conf` (блок `configs.nginx.file: ./default.conf`).

### DB

Файл `.env`: `DB_USER`, `DB_PASS`. Файл `users.sql` для инициализации (текст как в курсе: `grant all on database app...`, `\connect app`, `create table users...`).

Сервис `db`: образ `postgres`, переменные из `.env`, монтирование `users.sql` в `/docker-entrypoint-initdb.d/`. `docker compose up -d`, проверка таблицы через `docker exec`.

### Back

Приложение на Go (как в курсе): чтение строки подключения из `/run/secrets/connection_string`, выдача JSON пользователей. В `.env` добавить `DB_CONN="postgres://app:pass@db:5432/app?sslmode=disable"`.

`Dockerfile`: multi-stage `golang:1.21` → бинарь в `scratch`.

В `compose.yaml` сервис `back`: `build: .`, `depends_on: db`, `secrets.connection_string` из переменной окружения `DB_CONN`. Запуск: `docker compose up -d --build`. Вставка пользователя через `docker exec` (команда как в курсе). Убедиться в работе страницы.

### Volume

Проблема: при пересоздании контейнеров данные БД теряются (`docker compose rm -sf`). Решение: именованный том `db-data` для `/var/lib/postgresql/data` у сервиса `db`, блок `volumes: db-data:`. Проверить сохранность данных после remove/up.

### Network

Изоляция: две сети `front` и `db`; `front` только в `front`, `back` в обеих, `db` только в `db`. Проверка: с контейнера front недоступен хост `db` (`Could not resolve host` после применения конфигурации).

## Требования

- Vagrant-ВМ с Docker и Compose (или локально Docker Compose v2).
- Файлы проекта: `compose.yaml`, `index.html`, `default.conf`, `.env`, `users.sql`, `main.go`, `Dockerfile`.

## Шаги выполнения (предложенные)

1. Минимальный `compose.yaml` с nginx, `up -d`, проверка порта 8888.
2. Добавить front с кастомной страницей и конфигом через configs/volumes.
3. Добавить Postgres и инициализацию SQL.
4. Добавить сервис back со сборкой, secrets, проверкой E2E.
5. Ввести том для данных БД и проверить персистентность.
6. Разделить сети front/db и проверить изоляцию.

## Ожидаемый результат

- `curl` к localhost:8888 показывает кастомную страницу и данные из API back после добавления записей в БД.
- С томом `db-data` данные переживают пересоздание контейнеров.
- Front не достучится до `db` напрямую в финальной сетевой схеме.

## Команды (выделено для CLI-агента)

```bash
docker compose up -d
curl -s localhost:8888 | grep title
docker compose ls
docker compose ps
docker compose top
```

```bash
docker compose up -d
curl -s localhost:8888 | grep h2
```

```bash
docker exec -it app-db-1 su - postgres -c 'psql -U app app -c \d'
```

```bash
docker compose up -d --build
docker exec -it app-db-1 su - postgres -c "psql -U app app -c \"insert into users (name,email) values ('alex', 'alex@mail.ru');\""
```

```bash
docker compose rm -sf
docker compose up -d
```

```bash
docker exec app-front-1 curl -sS db:5432
docker compose up -d
docker exec app-front-1 curl -sS db:5432
```

Полные **`compose.yaml`** (все этапы), **`index.html`**, **`default.conf`**, **`main.go`**, **`Dockerfile`**, **`users.sql`**, **`.env`** — скопировать с [страницы курса](https://yudolevich.github.io/infra-course/06_practice_docker_compose.html).
