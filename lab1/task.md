# Vagrant

**Источник курса:** [01_practice_vagrant.html](https://yudolevich.github.io/infra-course/01_practice_vagrant.html)

## Задание

В данной практике рассмотрим базовое взаимодействие с [vagrant](https://www.vagrantup.com/).

### Install

#### VirtualBox

Для работы с vagrant нам понадобится провайдер, использующийся по умолчанию — [VirtualBox](https://www.virtualbox.org/). Скачать его можно со [страницы загрузок на официальном сайте](https://www.virtualbox.org/wiki/Downloads) или воспользовавшись пакетным менеджером для своей ОС.

#### Vagrant

Сам [vagrant](https://www.vagrantup.com/) можно [скачать и установить с официального сайта](https://developer.hashicorp.com/vagrant/downloads), где указаны инструкции для разных платформ.

После установки можно убедиться, что vagrant установлен: `vagrant -v` (в примере курса: `Vagrant 2.3.7`).

### Init

Для инициализации проекта воспользуемся командой `vagrant init`, которая создаст в текущей директории `Vagrantfile` с конфигурацией и комментариями к ней. Данная команда также позволяет задать имя бокса (образа виртуальной машины).

После этого в директории появится `Vagrantfile`. Если посмотреть содержимое, убрав комментарии и пустые строки, получится минимальная конфигурация с `config.vm.box = "ubuntu/lunar64"`.

### Box

К сожалению [HashiCorp](https://www.hashicorp.com/) на текущий момент ограничила доступ к своим ресурсам из России и автоматическое скачивание бокса при запуске не будет работать. Для установки бокса его [можно найти и вручную скачать с сайта](https://app.vagrantup.com/boxes/search). В данной практике используется [ubuntu/lunar64](https://app.vagrantup.com/ubuntu/boxes/lunar64), берётся версия для virtualbox (в примере: `20230829.0.0`).

После скачивания добавить бокс в локальную базу vagrant: `vagrant box add --name ubuntu/lunar64 <файл или URL>`.

Другой вариант — указать ссылку на бокс в `vagrant box add`; так как ссылка на app.vagrantup.com не прямая, необходимо получить прямую ссылку (в материале курса показано через `curl` с `-w '%{url_effective}'`), затем использовать её в `vagrant box add`.

### Usage

Для управления виртуальной машиной есть ряд команд:

- `vagrant up` — создаёт виртуальную машину согласно описанию в `Vagrantfile`
- `vagrant halt` — останавливает виртуальную машину
- `vagrant suspend` — отправляет в сон виртуальную машину
- `vagrant resume` — пробуждает от сна виртуальную машину
- `vagrant destroy` — полностью уничтожает виртуальную машину
- `vagrant reload` — перезапускает виртуальную машину, перечитывая `Vagrantfile`
- `vagrant status` — статус виртуальной машины
- `vagrant global-status` — глобальный статус по всем виртуальным машинам
- `vagrant port` — проброшенные порты виртуальной машины
- `vagrant ssh` — подключение к терминалу виртуальной машины через ssh

**Примечание:** по умолчанию vagrant синхронизирует текущий каталог с каталогом `/vagrant` внутри виртуальной машины.

**Примечание:** по умолчанию проброшен порт 22 с виртуальной машины на 2222 порту хоста для SSH.

Пробросить дополнительно порт 80 из ВМ на 8080 хоста: в `Vagrantfile` добавить строку `config.vm.network "forwarded_port", guest: 80, host: 8080` (в шаблоне от `vagrant init` она может быть в комментариях — достаточно раскомментировать). После изменений — `vagrant reload`.

### Provision

Через конфигурацию в `Vagrantfile` есть возможность подготовить виртуальную машину после запуска. Подготовить машину для запуска приложения: создать простое приложение, отвечающее на HTTP-запросы (в курсе пример на Go). Сохранить рядом с `Vagrantfile` (например `main.go`).

Добавить `shell` provisioner в `Vagrantfile`, который установит пакет `golang`, а также запустит приложение в фоне (`go run /vagrant/main.go &` в примере курса).

После `vagrant up` с пробросом 80→8080 проверка с хоста: `curl localhost:8080` — в примере ожидается ответ `Hello!`.

## Требования

- Провайдер по умолчанию: **VirtualBox** (или эквивалентный провайдер, совместимый с Vagrant).
- Установленный **Vagrant**.
- Бокс **`ubuntu/lunar64`** (или ручная установка бокса при ограничениях доступа к HashiCorp/Vagrant Cloud из региона).

## Шаги выполнения (предложенные)

1. Установить VirtualBox и Vagrant; проверить `vagrant -v`.
2. Создать каталог, выполнить `vagrant init ubuntu/lunar64`, при необходимости добавить бокс вручную (`vagrant box add`).
3. Выполнить `vagrant up`, `vagrant ssh`, проверить синхронизацию `/vagrant`, `vagrant status`, `vagrant port`.
4. Настроить проброс порта 80→8080, `vagrant reload`, снова проверить `vagrant port`.
5. (Раздел Provision) Создать HTTP-приложение, добавить provisioner с установкой Go и запуском приложения, `vagrant up` / `vagrant provision`, проверить `curl localhost:8080`.
6. По завершении: `vagrant halt` или `vagrant destroy`.

## Ожидаемый результат

- ВМ поднимается по `Vagrantfile`, доступ по `vagrant ssh`.
- В примере с provision и пробросом порта: запрос к `http://localhost:8080` с хоста возвращает тело ответа приложения (`Hello!` в примере курса).

## Команды (выделено для CLI-агента)

Установка (по ОС из материала курса):

```bash
# Windows (choco)
choco install virtualbox
choco install vagrant
```

```bash
# macOS (brew)
brew install --cask virtualbox
brew install hashicorp/tap/hashicorp-vagrant
```

```bash
# Ubuntu
sudo apt install virtualbox
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install vagrant
```

Инициализация и бокс:

```bash
mkdir vagrant
cd vagrant
vagrant init ubuntu/lunar64
grep -Ev '^\s*#|^$' Vagrantfile
```

```bash
curl -LO https://app.vagrantup.com/ubuntu/boxes/lunar64/versions/20230829.0.0/providers/virtualbox.box
vagrant box add --name ubuntu/lunar64 virtualbox.box
rm virtualbox.box
vagrant box list
```

```bash
curl -ILso /dev/null https://app.vagrantup.com/ubuntu/boxes/lunar64/versions/20230829.0.0/providers/virtualbox.box -w '%{url_effective}'
vagrant box add --name ubuntu/lunar64 "$(curl -ILso /dev/null https://app.vagrantup.com/ubuntu/boxes/lunar64/versions/20230829.0.0/providers/virtualbox.box -w '%{url_effective}')"
```

Жизненный цикл ВМ:

```bash
vagrant up
vagrant ssh
vagrant status
vagrant global-status
vagrant port
vagrant reload
vagrant halt
vagrant destroy
```

Пример `main.go` и создание файла (как в курсе):

```bash
cat <<'EOF' > main.go
package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("0.0.0.0:80", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello!\n"))
		}),
	)
}
EOF
```

Проверка приложения с хоста:

```bash
curl localhost:8080
```
