# Docker Volume/Network

**Источник курса:** [05_practice_docker_volume_network.html](https://yudolevich.github.io/infra-course/05_practice_docker_volume_network.html)

## Задание

В данном практическом занятии рассматривается работа с [томами (volume)](https://docs.docker.com/storage/volumes/) и конфигурациями [сети (network)](https://docs.docker.com/network/).

### Volume

#### Local

Стартовый `Vagrantfile` с одной машиной `storage`: Ubuntu lunar64, hostname `storage`, private_network dhcp, provision: каталог `/data`, nfs-server, `docker.io`, пользователь `vagrant` в группе `docker`.

- Монтирование каталога хоста/ВМ: `docker run -d -v /data:/usr/share/nginx/html -p 8888:80 --name nginx nginx`, запись в `/data/index.html`, `curl localhost:8888`.
- Именованный том: `docker volume create`, метки `--label`, `docker volume ls`, фильтр `-f`, `docker volume prune`.
- Имя тома в `-v html:/usr/share/nginx/html` — Docker создаёт том; **примечание курса:** при пустом томе на путь с файлами в образе данные из образа копируются в том.
- `docker volume inspect`, правка файлов в `Mountpoint`.
- Несколько монтирований и общий том между контейнерами; том переживает удаление контейнера.

#### Remote

Расширенный `Vagrantfile`: машины `storage`, `docker1`, `docker2` с docker и NFS-клиентом/сервером как в курсе.

На `docker1`/`docker2`: создать NFS-том и запустить nginx (в курсе `docker volume create -o type=nfs -o device=:/data -o o=addr=storage.local storage`).

Проверка с `storage`: `curl docker1.local:8888`, изменение `/data/index.html` отражается на обоих клиентах.

### Network

`docker network ls`.

#### Bridge

`docker network create br --subnet 10.0.0.0/24`, `docker network inspect br`.

На default bridge имя контейнера не резолвится; на пользовательской bridge — резолвинг имён и заданная подсеть. Пример с `alpine sleep inf`, `ping` между `first` и `second`.

`docker network connect` / `docker network disconnect` для подключения контейнера к нескольким сетям.

#### Host

`docker run ... --network host` — процессы слушают порты на адресе хоста ВМ.

#### IPVLAN

Объединение сетей Docker на нескольких ВМ через [драйвер ipvlan](https://docs.docker.com/network/drivers/ipvlan/): `parent` — интерфейс с адресом из сети Vagrant (в курсе для VirtualBox часто `192.168.56.0/24`, имя интерфейса через `ip -br a | awk '/192.168.56/{print $1}'` — пример `enp0s8`).

Создать `internal` сеть на `storage`, `docker1`, `docker2` с разными `--ip-range`, запустить nginx с томом `storage` и сетью `internal`, проверить IP контейнеров.

На `storage`: конфиг nginx upstream на IP контейнеров `docker1`/`docker2`, контейнер nginx с прокси и подключением к сети `internal`. Проверка `curl localhost:8888`; при остановке одной ВМ — запросы уходят на оставшуюся; при останове обеих — 502.

## Требования

- Vagrant с несколькими ВМ для разделов Remote и IPVLAN.
- Docker на соответствующих машинах; для NFS — настроенный экспорт `/data` как в курсе.

## Шаги выполнения (предложенные)

1. Local volume: bind mount, named volume, inspect, совместное использование тома.
2. Поднять три ВМ, настроить NFS-тома и проверить единое содержимое.
3. Bridge: создать сеть, сравнить DNS default vs custom, connect/disconnect.
4. Host network с nginx.
5. IPVLAN: создать сети, проверить связность и балансировку через nginx на storage.

## Ожидаемый результат

- Поведение томов и сетей совпадает с примерами вывода в курсе (`curl`, `ping`, ответы при остановке ВМ).

## Команды (выделено для CLI-агента)

```bash
docker run -d -v /data:/usr/share/nginx/html -p 8888:80 --name nginx nginx
docker volume create empty
docker volume create new --label test=true
docker volume ls
docker volume ls -f label=test=true
docker volume prune --filter label=test=true -af
docker volume inspect html
docker run -d -v html:/usr/share/nginx/html -p 8888:80 --name nginx nginx
```

```bash
docker network create br --subnet 10.0.0.0/24
docker network inspect br
docker run -d --network br --name first alpine sleep inf
docker run -d --network br --name second alpine sleep inf
docker network connect br first
docker network disconnect bridge first
```

```bash
docker run -d -v /data:/usr/share/nginx/html --network host --name nginx nginx
```

```bash
ip -br a | awk '/192.168.56/{print $1}'
docker network create -d ipvlan --subnet=10.1.1.0/24 --ip-range=10.1.1.0/28 -o parent=enp0s8 internal
```

Полные **`Vagrantfile`** фрагменты и блоки **`default.conf`** для upstream — с [страницы курса](https://yudolevich.github.io/infra-course/05_practice_docker_volume_network.html).
