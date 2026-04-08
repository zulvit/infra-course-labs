# Docker

**Источник курса:** [03_practice_docker.html](https://yudolevich.github.io/infra-course/03_practice_docker.html)

## Задание

В данном практическом занятии вспомнить основные возможности docker-клиента.

### Vagrant

Для работы с Docker вне зависимости от платформы можно использовать следующий `Vagrantfile`:

```ruby
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/lunar64"
  config.vm.provision "docker"
end
```

### Container lifecycle

#### Run

- Простой запуск: `docker run` с образом и командой (в курсе: `docker run python:3.11.5-alpine python --version`).
- Список запущенных контейнеров: `docker ps`; с остановленными: `docker ps -a`.
- Очистка остановленных: `docker container prune`.
- Автоудаление после выхода: `docker run --rm ...`.
- Интерактивный режим: `docker run --rm -it ...`.
- Фон: `docker run -d`, имя: `--name` (в курсе пример с `python -m http.server 8888`).

#### Exec

`docker exec` — выполнить команду внутри контейнера (в курсе: `docker exec test wget -qO- localhost:8888`).

#### Stop

`docker stop`, затем при необходимости `docker rm` для удаления контейнера.

#### Run Options

Пример HTTP-сервера на Python в файле `test.py` (слушает порт 8888, переменная окружения `FILE`) — разместить рядом с `Vagrantfile`, в ВМ путь `/vagrant/test.py`.

- Монтирование: `docker run -v /vagrant:/vagrant ...`
- Переменные и порты: `-e`, `-p` (в курсе пример с `FILE=/vagrant/Vagrantfile` и `curl` с хоста).

#### Logs / Ports / Stats

- `docker logs <container>`
- `docker port <container>`
- `docker top`, `docker stats`

### Images

Образы подтягиваются из registry (по умолчанию [hub.docker.com](https://hub.docker.com/)). Список локальных: `docker images`.

- `docker pull` (в курсе: `registry:2`, запуск локального registry на порту 5000).
- `docker tag`, `docker push` в локальный registry (`localhost:5000/...`).
- `docker save`, `docker load`.
- `docker build` с `Dockerfile` (пример: `FROM python:3.11.5-alpine`, `ADD test.py`, `CMD`).

### Remote

Docker-клиент по умолчанию общается с демоном через unix socket; можно использовать TCP. В курсе: изменить `Vagrantfile` — проброс порта 2375, `post_install_provision` для docker unit: добавить `-H tcp://0.0.0.0:2375`, перезапуск docker. На хосте: `export DOCKER_HOST=localhost:2375` и работа `docker` как с удалённым демоном.

## Требования

- Среда с **Docker** (в курсе — ВМ `ubuntu/lunar64` с provision `docker`).
- Доступ к загрузке образов из **Docker Hub** (или заранее подготовленные образы).

## Шаги выполнения (предложенные)

1. Поднять ВМ с Docker по примеру `Vagrantfile`.
2. Пройти сценарии `docker run` / `ps` / `prune` / `--rm` / `-it` / `-d` / `--name`.
3. Потренировать `docker exec`, `stop`, `rm`, `-f`.
4. Создать `test.py`, отработать `-v`, `-e`, `-p`, `logs`, `port`, `top`, `stats`.
5. Раздел Images: `pull`, локальный `registry`, `tag`/`push`, `save`/`load`, `build`.
6. (Опционально по курсу) Настроить удалённый доступ к демону через TCP и `DOCKER_HOST`.

## Ожидаемый результат

- Понимание жизненного цикла контейнера и базовых опций `docker run`.
- Работа с образами, локальным registry, сборкой через `Dockerfile`.
- В примере с `test.py` и пробросом порта ответ `curl` с хоста содержит содержимое выбранного файла.

## Команды (выделено для CLI-агента)

```bash
docker run python:3.11.5-alpine python --version
docker ps
docker ps -a
docker container prune
docker run --rm python:3.11.5-alpine python --version
docker run --rm -it python:3.11.5-alpine python
docker run --name test -d python:3.11.5-alpine python -m http.server 8888
docker exec test wget -qO- localhost:8888
docker stop test
docker rm test
docker run --rm python:3.11.5-alpine python --version
```

```bash
docker run -v /vagrant:/vagrant -d python:3.11.5-alpine python /vagrant/test.py
docker rm -f test
docker run -p 8888:8888 -v /vagrant:/vagrant -e FILE=/vagrant/Vagrantfile -d --name test python:3.11.5-alpine python /vagrant/test.py
curl localhost:8888
docker logs test
docker port test
docker top test
docker stats
```

```bash
docker images
docker pull registry:2
docker run -d -p 5000:5000 --name registry registry:2
docker tag python:3.11.5-alpine localhost:5000/python:3.11.5-alpine
docker push localhost:5000/python:3.11.5-alpine
docker save python:3.11.5-alpine -o python.tar
docker load -i python.tar
docker build -t test /vagrant/
docker rm -f test
docker run -p 8888:8888 -d --name test test
curl localhost:8888
```

```bash
export DOCKER_HOST=localhost:2375
docker run -d registry:2
docker ps
```

Текст **`test.py`** и полный **`Vagrantfile`** для TCP — взять с [страницы курса](https://yudolevich.github.io/infra-course/03_practice_docker.html).
