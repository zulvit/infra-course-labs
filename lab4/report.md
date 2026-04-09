# Отчёт

## Шаги выполнения
1. Подготовлены каталоги `assets/screenshots`, `artifacts/logs`, `work`.
2. Для user-space выполнения использованы уже установленные в ходе `lab3` инструменты:
   - `udocker` как rootless/containerless runtime;
   - `martizih/kaniko:latest` как builder для Dockerfile без Docker daemon.
3. Подготовлены примеры:
   - [work/hello/Dockerfile](/home/savva/Desktop/devops/infra-course-labs/lab4/work/hello/Dockerfile)
   - [work/main/Dockerfile](/home/savva/Desktop/devops/infra-course-labs/lab4/work/main/Dockerfile)
   - [work/main/main.py](/home/savva/Desktop/devops/infra-course-labs/lab4/work/main/main.py)
   - [work/multi/Dockerfile](/home/savva/Desktop/devops/infra-course-labs/lab4/work/multi/Dockerfile)
   - [work/multi/main.go](/home/savva/Desktop/devops/infra-course-labs/lab4/work/multi/main.go)
4. Выполнена успешная сборка минимального образа `hello` через `kaniko` с сохранением артефакта `hello.tar`.
5. Поднят локальный `registry:2` через `udocker`, проверен HTTP API `/v2/`.
6. Выполнен прямой push образа `hello:0.1` в локальный registry через `kaniko`, после чего проверены `_catalog` и `tags/list`.
7. Попытки собрать более сложные образы `main` и `multi` завершились предсказуемой ошибкой user-space среды: на этапе unpack rootfs `kaniko` получил `operation not permitted` на `chown`.
8. Попытка импортировать `kaniko` tar в `udocker` для runtime-проверки показала ещё одно ограничение: rootfs импортируется, но metadata образа (`CMD`, `ENTRYPOINT`) теряется, поэтому корректный запуск собранных образов через `udocker import` не воспроизводится.

## Команды
- `udocker pull martizih/kaniko:latest`
- `udocker run --entrypoint="" martizih/kaniko:latest /kaniko/executor ... --tar-path=... --no-push --force`
- `udocker pull registry:2`
- `udocker run registry:2`
- `wget -qO- http://127.0.0.1:5000/v2/`
- `udocker run --entrypoint="" martizih/kaniko:latest /kaniko/executor ... --destination=127.0.0.1:5000/hello:0.1 --insecure ...`
- `wget -qO- http://127.0.0.1:5000/v2/_catalog`
- `wget -qO- http://127.0.0.1:5000/v2/hello/tags/list`

## Вывод
- Удалось воспроизвести часть лабораторной без системного Docker:
  - сборка простого Dockerfile без unpack rootfs;
  - запуск локального registry;
  - push образа в registry и проверка Registry HTTP API.
- Подтверждённые результаты:
  - `registry /v2/` отвечает `{}`;
  - `_catalog` возвращает `{"repositories":["hello"]}`;
  - `tags/list` возвращает `{"name":"hello","tags":["0.1"]}`.

Актуализация и ограничения:
- В исходном задании предполагается полноценный Docker daemon.
- В текущей среде без root сборка через `kaniko` возможна только для простых случаев, не требующих unpack rootfs с операциями `chown`.
- Для `python:3.11-alpine` и `golang:1.21` сборка оборвалась на ошибках вида:
  - `failed to get filesystem from image: chown /bin: operation not permitted`
  - `failed to get filesystem from image: chown /boot: operation not permitted`
- Импорт `kaniko` tar в `udocker` даёт только файловую систему слоя, но не сохраняет runtime metadata образа, поэтому разделы с корректным `docker run` для собранных образов в таком окружении не воспроизводятся.
- Следовательно, `lab4` выполнена частично: build/registry-сценарии подтверждены, а runtime-проверки сложных образов честно задокументированы как заблокированные ограничениями среды.

## Скриншоты
- Вместо графических скриншотов сохранены текстовые логи:
  - [lab4_kaniko_probe.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4_kaniko_probe.typescript)
  - [session_lab4.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/session_lab4.typescript)
  - [registry_probe.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/registry_probe.typescript)
  - [registry_probe.out](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/registry_probe.out)
  - [registry_push.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/registry_push.typescript)
  - [hello.tar](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/hello.tar)
  - [main_build.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/main_build.typescript)
  - [multi_build.typescript](/home/savva/Desktop/devops/infra-course-labs/lab4/artifacts/logs/multi_build.typescript)
