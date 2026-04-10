# Лабораторная работа 4

## Тема

Сборка контейнерных образов и работа с локальным registry

## Цель работы

Изучить сборку Docker-образов по `Dockerfile`, запуск собранных контейнеров, применение директив `ARG`, `ENV`, `RUN`, `COPY`, `ADD`, `EXPOSE`, `LABEL`, `ENTRYPOINT`, `CMD`, сборку multi-stage образа и публикацию образа в локальный registry.

## Ход выполнения

### 1. Подготовка среды

Для выполнения лабораторной использован локальный Docker daemon.  
В рабочем каталоге были подготовлены три примера:

- `hello` — минимальный образ на `busybox`;
- `main` — Python/FastAPI приложение с демонстрацией основных директив Dockerfile;
- `multi` — multi-stage сборка бинарного Go-приложения.

### 2. Сборка и проверка минимального образа

Образ `lab4/hello:0.1` был собран из каталога `lab4/work/hello`.

Результат:

- сборка завершилась успешно;
- команда `docker run --rm lab4/hello:0.1` вернула строку `Hello from hello image`.

### 3. Сборка и проверка образа `main`

Из каталога `lab4/work/main` был собран образ `lab4/main:0.1` с аргументом `FILE=/example.html`.

В ходе сборки и проверки были подтверждены:

- `ARG FILE` и `ENV FILE="${FILE}"`;
- `RUN pip install --no-cache-dir fastapi "uvicorn[standard]"`;
- `COPY --chmod=555 main.py /app/main.py`;
- `ADD https://example.com /example.html`;
- `EXPOSE 8888`;
- `LABEL version="0.1"`;
- `ENTRYPOINT ["uvicorn", "main:app"]`;
- `CMD ["--host=0.0.0.0", "--port=8888"]`.

После запуска контейнера на `127.0.0.1:8888` HTTP-запрос вернул страницу `Example Domain`, что подтвердило корректную работу приложения и чтение файла `/example.html`.

Дополнительные проверки `docker image inspect` показали:

- `ExposedPorts` → `{"8888/tcp":{}}`
- `Labels` → `{"version":"0.1"}`
- `Entrypoint/Cmd` → `["uvicorn","main:app"] ["--host=0.0.0.0","--port=8888"]`

### 4. Сборка и проверка multi-stage образа

Из каталога `lab4/work/multi` был собран образ `lab4/multi:0.1`.

Результат:

- сборка прошла успешно;
- контейнер вернул строку `Hello from multi-stage image`;
- `docker history` показал, что итоговый образ содержит только скопированный бинарник и `ENTRYPOINT`, что соответствует ожидаемому multi-stage сценарию;
- итоговый размер образа составил около `1.12MB`.

### 5. Публикация образа в локальный registry

Локальный реестр был запущен командой:

```text
docker run -d -p 5000:5000 --name lab4-registry registry:2
```

После этого образ `lab4/hello:0.1` был опубликован как `127.0.0.1:5000/hello:0.1`.

Подтверждённые результаты:

- `/v2/` отвечает корректно;
- `/_catalog` возвращает `{"repositories":["hello"]}`;
- `/hello/tags/list` возвращает `{"name":"hello","tags":["0.1"]}`;
- `docker push` завершился с digest `sha256:8788245b202d9478967d9c25d7f17978c04d81e2e330e470851465e746b53320`.

Для проверки манифеста использован заголовок:

```text
Accept: application/vnd.oci.image.index.v1+json
```

Это связано с тем, что текущий Docker push опубликовал OCI index, а не старый docker v2 manifest.

## Результаты

В ходе лабораторной работы были успешно подтверждены:

- сборка минимального образа;
- сборка и запуск образа с Python/FastAPI приложением;
- работа директив `ARG`, `ENV`, `RUN`, `COPY`, `ADD`, `EXPOSE`, `LABEL`, `ENTRYPOINT`, `CMD`;
- сборка и запуск multi-stage образа;
- запуск локального `registry:2`;
- публикация образа в registry;
- проверка Registry HTTP API.

## Вывод

Лабораторная работа 4 выполнена полностью.  
Все ключевые сценарии из задания подтверждены на практике: образы собираются и запускаются, multi-stage сборка работает корректно, локальный registry принимает опубликованный образ и отвечает через HTTP API.
