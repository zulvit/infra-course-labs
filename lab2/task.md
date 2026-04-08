# Vagrant Multi-Node

**Источник курса:** [02_practice_vagrant.html](https://yudolevich.github.io/infra-course/02_practice_vagrant.html)

## Задание

В данной практике рассмотрим такие функции [vagrant](https://www.vagrantup.com/), как создание собственного [box](https://developer.hashicorp.com/vagrant/docs/boxes) из развёрнутой машины, а также создание нескольких машин в одном [Vagrantfile](https://developer.hashicorp.com/vagrant/docs/vagrantfile).

В качестве лабораторного стенда реализовать простую схему развёртывания из трёх виртуальных машин:

- **front** — отдаёт статическую HTML-страницу;
- **back** — приложение с бизнес-логикой;
- **db** — база данных.

### Package

Для удобства развёртывания на основе базового бокса создать боксы с нужной конфигурацией, чтобы не готовить окружение при каждом запуске. Для этого используется команда `vagrant package`, которая сохраняет бокс из текущей запущенной ВМ на файловую систему.

#### Front

Базовый бокс `ubuntu/lunar64`, `Vagrantfile` с provision: установка `nginx` для отдачи статической страницы. После `vagrant up` — сохранить приватный ключ для SSH (путь в `IdentityFile` из `vagrant ssh-config`); в курсе показано копирование в `./key`.

Далее: `vagrant package --output front.box`, `vagrant box add --name front front.box`, `vagrant destroy -f`, удалить `front.box`.

#### Back

Для back: `ubuntu/lunar64`, provision — установка `golang`. Те же действия: `vagrant up`, `vagrant package --output back.box`, `vagrant box add --name back back.box`, `vagrant destroy -f`, `rm back.box`.

#### DB

Для db: `ubuntu/lunar64`, установка `postgresql`, конфигурация для удалённого подключения: в курсе добавляются `listen_addresses = '*'` и строка `host all all 0.0.0.0/0 trust` в соответствующие файлы PostgreSQL (пути как в примере для версии 15 на Ubuntu). Аналогично: package → `db.box` → `vagrant box add --name db db.box` → destroy → удалить файл бокса.

### Multi-Node

После подготовки в `vagrant box list` должны появиться боксы `front`, `back`, `db` (и базовый при необходимости).

#### Code

**Front:** файл `index.html` — статическая страница с JS, запрашивающим данные (в примере курса запрос на `http://localhost:8889` — при переносе на multi-node конфигурация меняется на схему из `Vagrantfile`).

**Back:** `main.go` — HTTP-сервис на Go с подключением к PostgreSQL (`github.com/jackc/pgx/v5`), строка подключения в примере: `postgres://app:pass@db:5432/app?sslmode=disable`, отдача JSON со списком пользователей.

**DB:** `users.sql` — создание БД, пользователя, таблицы `users`, выдача прав (текст — как в материале курса).

#### Vagrantfile

Описать три машины `db`, `front`, `back` с `private_network` (в примере IP: db `192.168.56.30`, front `192.168.56.10`, back `192.168.56.20`), provision с копированием файлов из `/vagrant`, записями в `/etc/hosts`, на back — сборка и запуск Go-приложения. Использовать `config.ssh.private_key_path = "key"` (ключ из этапа package).

#### UP

В каталоге: `index.html`, `key`, `main.go`, `users.sql`, `Vagrantfile`. Команда `vagrant up`.

Проверка: открыть в браузере `http://localhost:8888` (проброс с front), убедиться, что страница работает.

Добавить пользователей в БД (в курсе: `vagrant ssh db -c 'sudo -u postgres psql'`, затем SQL `INSERT` в таблицу `users`). Убедиться, что данные отображаются на фронтенде.

После работы: `vagrant destroy -f`.

## Требования

- Vagrant + VirtualBox (как в предыдущей практике).
- Подготовленные локальные боксы **`front`**, **`back`**, **`db`** по инструкции раздела Package.
- Файл **`key`** — SSH-ключ для доступа к ВМ (как в курсе).

## Шаги выполнения (предложенные)

1. Последовательно собрать и добавить боксы front, back, db через `vagrant package` / `vagrant box add`.
2. Подготовить `index.html`, `main.go`, `users.sql` и общий `Vagrantfile` для трёх узлов.
3. Поместить `key` в каталог проекта, выполнить `vagrant up`.
4. Проверить UI на `http://localhost:8888`, добавить записи в БД, убедиться в обновлении таблицы на странице.
5. `vagrant destroy -f`.

## Ожидаемый результат

- Три ВМ поднимаются одной конфигурацией; front доступен с хоста на порту 8888.
- После вставки пользователей в БД таблица на веб-странице показывает актуальные данные.

## Команды (выделено для CLI-агента)

```bash
vagrant up
vagrant ssh-config | grep IdentityFile
vagrant ssh-config | awk '/IdentityFile/{print $2}' | xargs -I{} cp {} ./key
```

```bash
vagrant package --output front.box
vagrant box add --name front front.box
vagrant destroy -f
rm front.box
```

```bash
vagrant package --output back.box
vagrant box add --name back back.box
vagrant destroy -f
rm back.box
```

```bash
vagrant package --output db.box
vagrant box add --name db db.box
vagrant destroy -f
rm db.box
```

```bash
vagrant box list
vagrant up
vagrant ssh db -c 'sudo -u postgres psql'
vagrant destroy -f
```

Код `index.html`, `main.go`, `users.sql` и полный `Vagrantfile` — перенести дословно из [страницы курса](https://yudolevich.github.io/infra-course/02_practice_vagrant.html) (разделы Code и Vagrantfile).
