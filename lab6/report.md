# Отчёт

## Шаги выполнения
1. После получения `sudo` установлен и запущен локальный Docker Compose v2, поэтому лабораторная выполнялась напрямую на текущей машине без `Vagrant`.
2. Подготовлены файлы проекта в каталоге [work](/home/savva/Desktop/devops/infra-course-labs/lab6/work):
   - [compose.yaml](/home/savva/Desktop/devops/infra-course-labs/lab6/work/compose.yaml)
   - [index.html](/home/savva/Desktop/devops/infra-course-labs/lab6/work/index.html)
   - [default.conf](/home/savva/Desktop/devops/infra-course-labs/lab6/work/default.conf)
   - [.env](/home/savva/Desktop/devops/infra-course-labs/lab6/work/.env)
   - [users.sql](/home/savva/Desktop/devops/infra-course-labs/lab6/work/users.sql)
   - [main.go](/home/savva/Desktop/devops/infra-course-labs/lab6/work/main.go)
   - [Dockerfile](/home/savva/Desktop/devops/infra-course-labs/lab6/work/Dockerfile)
3. Выполнен запуск стека `front + back + db` через [work/run.sh](/home/savva/Desktop/devops/infra-course-labs/lab6/work/run.sh):
   - `docker compose up -d --build`
   - `docker compose ls`
   - `docker compose ps`
   - `docker compose top`
4. Проверен frontend:
   - по `http://127.0.0.1:18888` открывается кастомная страница с заголовком `Lab6 Compose`;
   - запрос `http://127.0.0.1:18888/back` возвращает JSON из backend.
5. Проверена работа БД:
   - после запуска API вернуло начальные записи `john` и `mary`;
   - через `docker exec` в PostgreSQL добавлен пользователь `alex`;
   - повторный запрос к `/back` вернул уже три записи.
6. Проверена персистентность:
   - выполнены `docker compose rm -sf` и повторный `docker compose up -d`;
   - запись `alex` сохранилась, значит именованный том `db-data` работает корректно.
7. Проверена изоляция сетей:
   - `front` подключён только к сети `front`;
   - `db` подключён только к сети `db`;
   - `back` подключён к обеим;
   - из контейнера `front` имя `db` не резолвится, что соответствует ожидаемому результату финальной схемы.

## Команды
- `sudo -n docker compose --project-name lab6app up -d --build`
- `sudo -n docker compose --project-name lab6app ls`
- `sudo -n docker compose --project-name lab6app ps`
- `sudo -n docker compose --project-name lab6app top`
- `wget -qO- http://127.0.0.1:18888`
- `wget -qO- http://127.0.0.1:18888/back`
- `sudo -n docker exec lab6app-db-1 psql -U app -d app -c "insert into users (name, email) values ('alex', 'alex@mail.ru');"`
- `sudo -n docker compose --project-name lab6app rm -sf`
- `sudo -n docker compose --project-name lab6app up -d`
- `sudo -n docker exec lab6app-front-1 sh -lc 'wget -T 2 -qO- http://db:5432'`

## Вывод
- `docker compose` поднял три связанных сервиса: `front`, `back`, `db`.
- Кастомная страница доступна на `127.0.0.1:18888`; вместо учебного `8888` использован `18888`, чтобы не конфликтовать с занятыми портами хоста.
- Backend корректно читает строку подключения из секрета и отдаёт JSON со списком пользователей.
- После вставки записи в PostgreSQL данные сразу видны через frontend/backend цепочку.
- После пересоздания контейнеров запись `alex` сохраняется, что подтверждает работу тома `db-data`.
- Финальная схема сетей работает как ожидается: из `front` прямой доступ к `db` отсутствует, ошибка в логе `wget: bad address 'db:5432'`.

Адаптации:
- Вместо `Vagrant` использован локальный Docker Compose v2, что прямо допускается в задании.
- Образы выбраны в актуальных вариантах `nginx:1.25-alpine`, `postgres:16-alpine`; это эквивалентные современные замены.

## Скриншоты
- Вместо графических скриншотов сохранены текстовые логи:
  - [session_lab6.typescript](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/session_lab6.typescript)
  - [front_index.out](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/front_index.out)
  - [back_initial.json](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/back_initial.json)
  - [back_after_insert.json](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/back_after_insert.json)
  - [back_after_recreate.json](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/back_after_recreate.json)
  - [front_to_db.err](/home/savva/Desktop/devops/infra-course-labs/lab6/artifacts/logs/front_to_db.err)
