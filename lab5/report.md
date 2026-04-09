# Отчёт

## Шаги выполнения
1. После получения `sudo` установлен нативный Docker и Compose:
   - `uidmap`
   - `docker.io`
   - `docker-compose-v2`
2. Подготовлен сценарий [work/run.sh](/home/savva/Desktop/devops/infra-course-labs/lab5/work/run.sh) и исходный HTML [work/index.html](/home/savva/Desktop/devops/infra-course-labs/lab5/work/index.html).
3. Отработан bind mount:
   - каталог `lab5/work/data` смонтирован в `nginx`;
   - проверен ответ по `http://127.0.0.1:18088`.
4. Отработаны named volumes:
   - созданы `lab5-empty`, `lab5-new`, `lab5-html`;
   - показан фильтр по label;
   - содержимое volume изменено через `Mountpoint`;
   - проверен ответ по `http://127.0.0.1:18089`.
5. Подтверждено совместное использование volume:
   - контейнер `lab5-writer` записал файл в `lab5-html`;
   - контейнер `lab5-reader` прочитал тот же файл.
6. Отработаны network-сценарии:
   - на default bridge имя контейнера не резолвится;
   - создана пользовательская bridge сеть `lab5-br` с подсетью `10.0.0.0/24`;
   - контейнеры `lab5-first` и `lab5-second` успешно пингуются по имени;
   - проверены `docker network connect` и `disconnect`.
7. Отработан `host` network:
   - контейнер с `python -m http.server 18891` доступен напрямую на адресе хоста.
8. Multi-host разделы `Remote` и полноценный `IPVLAN` из курса не воспроизводились, так как текущая лабораторная среда не содержит нескольких ВМ/хостов и NFS-топологии как в исходном материале.

## Команды
- `sudo -n docker pull nginx:alpine`
- `sudo -n docker pull alpine`
- `sudo -n docker pull python:3.11-alpine`
- `sudo -n docker run -d -v ... -p 18088:80 --name lab5-nginx-bind nginx:alpine`
- `sudo -n docker volume create ...`
- `sudo -n docker volume inspect lab5-html`
- `sudo -n docker run -d --name lab5-default-first alpine sleep inf`
- `sudo -n docker network create lab5-br --subnet 10.0.0.0/24`
- `sudo -n docker exec lab5-first ping -c 1 lab5-second`
- `sudo -n docker run -d --network host --name lab5-host-http python:3.11-alpine python -m http.server 18891`

## Вывод
- Bind mount работает: ответ nginx содержит `lab5 bind mount`.
- Named volume работает: ответ nginx содержит `lab5 named volume`.
- Один и тот же volume успешно разделяется между контейнерами, файл `shared.txt` читается из второго контейнера.
- На default `bridge` DNS-резолвинг имён контейнеров отсутствует (`ping: bad address`).
- На пользовательской сети `lab5-br` DNS-резолвинг имён работает, ping успешен.
- В `host` network сервис контейнера доступен напрямую на порту хоста без `-p`.

Адаптации:
- Вместо портов `8888/8889` использованы `18088/18089`, потому что низкие учебные порты уже были заняты в текущей среде.
- Разделы с NFS и multi-host `ipvlan` не выполнялись, поскольку для них нужна отдельная топология из нескольких машин, которой здесь нет.

## Скриншоты
- Вместо графических скриншотов сохранены текстовые логи:
  - [session_lab5.typescript](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/session_lab5.typescript)
  - [bind_mount_response.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/bind_mount_response.out)
  - [named_volume_response.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/named_volume_response.out)
  - [shared_volume.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/shared_volume.out)
  - [custom_bridge_ping.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/custom_bridge_ping.out)
  - [host_network_response.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/host_network_response.out)
  - [step01_env_note.out](/home/savva/Desktop/devops/infra-course-labs/lab5/artifacts/logs/step01_env_note.out)
