# Отчёт по лабораторной работе №4

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Собрать несколько Docker-образов с разной структурой Dockerfile, запустить приложения, посмотреть метаданные образов, поднять локальный registry и проверить API реестра.

## Ход работы

В каталоге `work` подготовлены образы `hello`, `main` (Python/uvicorn) и multi-stage `multi`. Сборка и запуск:

```powershell
docker build -t lab4/hello:0.1 ./work/hello
docker run --rm lab4/hello:0.1
docker build --build-arg FILE=/example.html -t lab4/main:0.1 ./work/main
docker run -d -p 8888:8888 --name lab4-main lab4/main:0.1
curl.exe -s http://127.0.0.1:8888/
docker build -t lab4/multi:0.1 ./work/multi
docker run --rm lab4/multi:0.1
```

Registry:

```powershell
docker run -d -p 5000:5000 --name lab4-registry registry:2
docker tag lab4/hello:0.1 127.0.0.1:5000/hello:0.1
docker push 127.0.0.1:5000/hello:0.1
curl.exe -s http://127.0.0.1:5000/v2/_catalog
```

Подробный вывод по шагам — в `artifacts/logs/step01_hello_build.log` … `step26_images_sizes.log`.

## Результаты (фрагменты вывода)

Запуск финального образа multi-stage:

```text
Hello from multi-stage image
```

Каталог репозиториев в registry:

```json
{"repositories":["hello"]}
```

Для запроса манифеста при необходимости указывался заголовок `Accept` под OCI index — иначе registry отвечал не тем типом (это видно в логах шага с manifest).

## Вывод

Все три образа собираются и запускаются, у `main` проверены expose/labels/entrypoint через `docker image inspect`. Локальный registry принимает push, список тегов и каталог репозиториев по HTTP API совпадают с ожиданиями.
