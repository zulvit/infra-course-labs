# Отчёт по лабораторной работе №1

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Поднять виртуальную машину в VirtualBox через Vagrant, развернуть простое HTTP-приложение на Go и убедиться, что оно отвечает с хоста по проброшенному порту.

## Ход работы

Собраны `vagrant/Vagrantfile` и `vagrant/main.go`. Приложение ставится в `/usr/local/bin/hello-vagrant` и запускается unit-файлом systemd `hello-vagrant.service`, чтобы сервис не зависел от интерактивной сессии после provision.

Проверка конфигурации:

```powershell
vagrant validate
```

Запуск ВМ:

```powershell
vagrant up
```

Проверка изнутри гостя (ядро, сервис, локальный HTTP):

```powershell
vagrant ssh -c "uname -a && systemctl status hello-vagrant.service --no-pager && curl -s http://127.0.0.1"
```

Проброс порта и запрос с хоста:

```powershell
vagrant port
curl.exe -s http://127.0.0.1:8080
```

После работы:

```powershell
vagrant destroy -f
```

## Результаты (фрагменты вывода)

Команда `vagrant port`:

```text
80 (guest) => 8080 (host)
```

Запрос с хоста:

```text
Hello!
```

Полные логи шагов лежат в каталоге `artifacts/logs` (например `step25_vagrant_ssh_checks_after_fix.log`, `step27_curl_localhost_8080_after_fix.log`).

## Замечания

В методичке мог быть указан образ `ubuntu/lunar64`; он уже неактуален, использован LTS `ubuntu/jammy64`. Сервис через systemd надёжнее, чем оставлять процесс только в shell provisioner.

## Вывод

ВМ поднимается, приложение в автозапуске отдаёт ответ `Hello!`, с хоста страница открывается на `http://127.0.0.1:8080`. Стенд по завершении снят командой `vagrant destroy -f`.
