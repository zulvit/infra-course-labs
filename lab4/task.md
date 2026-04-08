# Docker Images

**Источник курса:** [04_practice_docker_images.html](https://yudolevich.github.io/infra-course/04_practice_docker_images.html)

## Задание

В данном практическом занятии рассматриваются операции с docker-образами: директивы Dockerfile, сборка, работа с registry.

### Vagrant

```ruby
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/lunar64"
  config.vm.provision "docker"
end
```

### Dockerfile

[Директивы Dockerfile](https://docs.docker.com/engine/reference/builder/). Во всех примерах `Dockerfile` в директории проекта с `Vagrantfile`, внутри ВМ доступен как `/vagrant/Dockerfile`.

#### Hello

Минимальный образ: `FROM`, `CMD`. Сборка: `docker build -t hello /vagrant/`. Запуск: `docker run --rm hello:latest`. Информация: `docker image inspect hello:latest`.

#### Copy/Add

Скрипт `main.py` (HTTP на 8888, переменная `FILE`) рядом с Dockerfile.

Пример Dockerfile с `COPY main.py /main.py` и `ADD https://example.com /example.html`, `CMD ["python", "main.py"]`.

**Примечание курса:** отличие `COPY` от `ADD` — у `ADD` загрузка по URL и распаковка tar; рекомендуется по возможности `COPY`, а загрузки/распаковку делать в `RUN`.

Вариант с правами: `COPY --chmod=555 main.py /main.py`, `CMD ["/main.py"]`.

#### Env/Arg

`ARG FILE`, `ENV FILE="${FILE}"`, передача: `docker build -t main --build-arg FILE=/example.html /vagrant/`.

#### Run

Версия `main.py` на FastAPI; в Dockerfile: `RUN pip install fastapi "uvicorn[standard]"`, `CMD` с uvicorn.

**Cache:** порядок слоёв важен; установку зависимостей разместить до копирования часто меняющихся файлов (пример перестановки в курсе).

#### Expose

`EXPOSE 8888` — смотреть в `docker image inspect` и `docker ps`.

#### Label

`LABEL version="0.1"` — смотреть через `docker image inspect ... --format='{{json .Config.Labels}}'`.

#### Cmd/Entrypoint

Разделение: `ENTRYPOINT ["uvicorn", "main:app"]`, `CMD ["--host=0.0.0.0", "--port=8888"]`. Переопределение при `docker run` и через `--entrypoint`.

#### Multi-Stage

Сборка в `golang:1.21`, итоговый бинарник в `FROM scratch` (в курсе пример с `COPY --from=build` и встроенным `main.go` через heredoc в Dockerfile).

### Registry

Локальный registry: `docker run -d -p 5000:5000 --name registry registry:2`, проверка `curl localhost:5000/v2/` → `{}`.

`docker tag`, `docker push` на `localhost:5000/...`.

**Catalog:** `curl localhost:5000/v2/_catalog`

**Tags:** `curl localhost:5000/v2/hello/tags/list`

**Manifest:** `curl localhost:5000/v2/hello/manifests/0.1` и с заголовком `Accept: application/vnd.docker.distribution.manifest.v2+json`.

## Требования

- Docker в ВМ (как в разделе Vagrant курса).
- Сетевой доступ для `docker build`/`pull` (в т.ч. `example.com` для примеров `ADD` в курсе).

## Шаги выполнения (предложенные)

1. Пройти примеры Hello → COPY/ADD → ARG/ENV → RUN → кэш слоёв → EXPOSE → LABEL → CMD/ENTRYPOINT.
2. Собрать multi-stage образ `hello`, проверить размер и слои.
3. Поднять локальный `registry:2`, выполнить `tag`/`push`, запросы к API каталога/тегов/манифеста.

## Ожидаемый результат

- Собранные образы ведут себя как в пошаговых примерах курса (`docker run`, `curl` к приложению на 8888).
- Локальный registry отвечает на `/v2/`, образы пушатся и видны в `_catalog`.

## Команды (выделено для CLI-агента)

```bash
docker build -t hello /vagrant/
docker run --rm hello:latest
docker image inspect hello:latest
```

```bash
docker build -t main /vagrant/
docker run -d -p 8888:8888 --name main main:latest
curl localhost:8888
docker build -t main --build-arg FILE=/example.html /vagrant/
docker rm -f main
docker run -d -p 8888:8888 --name main main:latest
curl -s localhost:8888 | grep title
```

```bash
docker build -t main /vagrant/
docker image inspect main:latest --format='{{json .Config.ExposedPorts}}'
docker rm -f main
docker run -d --name main main:latest
docker ps
```

```bash
docker build -t hello /vagrant/
docker run --rm hello:latest
docker image inspect hello:latest --format='{{json .RootFS}}'
docker images hello
```

```bash
docker run -d -p 5000:5000 --name registry registry:2
curl localhost:5000/v2/
docker tag hello:latest localhost:5000/hello
docker push localhost:5000/hello
curl localhost:5000/v2/_catalog
curl localhost:5000/v2/hello/tags/list
```

Полные тексты **Dockerfile**, **main.py** (все варианты из занятия) — скопировать с [страницы курса](https://yudolevich.github.io/infra-course/04_practice_docker_images.html).
