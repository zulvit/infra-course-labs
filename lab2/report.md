# Отчёт по лабораторной работе №2

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Собрать три локальных box-образа (front, back, db), поднять многоузловой стенд и связать nginx, backend на Go и PostgreSQL.

## Ход работы

Подготовлены `build/front/Vagrantfile`, `build/back/Vagrantfile`, `build/db/Vagrantfile`, общий `Vagrantfile`, статика `index.html`, backend `main.go`, скрипты `users.sql`, `insert_users.sql`, `select_users.sql`. Образы собирались через `vagrant package`, добавлялись в локальный каталог `vagrant box add`, затем поднимался стенд `vagrant up`.

Проверка списка box:

```powershell
vagrant box list
```

Проверка backend и фронта с хоста:

```powershell
curl.exe -s http://127.0.0.1:8889
curl.exe -s http://127.0.0.1:8888
```

Работа с БД через SSH к ВМ `db` (порт смотреть в `vagrant ssh-config db`, в моём прогоне был проброс на 2202):

```powershell
ssh.exe -i key -p 2202 vagrant@127.0.0.1 "sudo -u postgres psql -d app -f /vagrant/insert_users.sql"
ssh.exe -i key -p 2202 vagrant@127.0.0.1 "sudo -u postgres psql -d app -c \"select * from users;\""
```

Завершение:

```powershell
vagrant destroy -f
```

## Результаты (фрагменты вывода)

Ответ API после вставки строк:

```json
[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]
```

Фрагмент вывода `psql`:

```text
 id | name
----+------
  1 | Alice
  2 | Bob
```

Подробные логи: `artifacts/logs/step21_multinode_up.log`, `step43_insert_users.log`, `step44_select_users_after_insert.log`, `step46_curl_host_8888_after_insert.log` и др.

## Замечания

Для jammy вместо устаревшего lunar. На Windows у файла ключа нужны нормальные права ACL, иначе OpenSSH ругается. Unit для backend без `User=vagrant`, если слушаем 80 от root — иначе не хватает прав на порт.

## Вывод

Стенд из трёх ВМ работает: фронт отдаёт страницу, backend ходит в Postgres, после вставки данных в таблице и в JSON совпадают имена Alice и Bob.
