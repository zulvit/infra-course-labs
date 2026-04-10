# Отчёт по лабораторной работе №7

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Поднять три ВМ (bastion, node1, node2), настроить доступ с bastion по SSH к узлам, выполнить ad-hoc команды Ansible и плейбук для установки nginx, конфигов и шаблонов.

## Ход работы

В репозитории: `Vagrantfile`, `hosts`, `playbook.yaml`, `files/default`, `templates/index.html.j2`.

Основные шаги:

```powershell
vagrant validate
vagrant up
```

На bastion: положить ключ для доступа к узлам (путь к ключу смотреть в `vagrant ssh-config bastion`), затем:

```bash
ansible -i /vagrant/hosts -m ping all
ansible -i /vagrant/hosts -m shell -a 'uname -a' nodes
ansible-playbook -i /vagrant/hosts /vagrant/playbook.yaml
```

Повторный запуск плейбука — проверка идемпотентности (`changed=0`). HTTP-проверки с bastion:

```bash
curl -s node1.local | grep title
curl node1.local/node1only
```

Цепочка команд и ошибки VirtualBox/Vagrant при первых попытках зафиксированы в `artifacts/logs/step01_*.log` … `step37_*.log`.

## Результаты (фрагменты вывода)

Ping всех хостов:

```text
bastion.local | SUCCESS => {"ping": "pong"}
node1.local   | SUCCESS => {"ping": "pong"}
node2.local   | SUCCESS => {"ping": "pong"}
```

Отдельная страница только на node1:

```text
node1 only
```

## Замечания

Сначала упёрлись в конфликт host-only DHCP у VirtualBox; вывели сеть на статические адреса в одном сегменте (например 192.168.57.x). Для сценария с общим ключом на bastion у Vagrant отключена подмена ключа (`ssh.insert_key = false`). Образ ВМ — jammy вместо устаревшего lunar из старых методичек.

## Вывод

Стенд работает, Ansible с bastion доходит до узлов, плейбук накатывает nginx и файлы, повторный прогон не ломает конфигурацию. HTTP-заголовки и ответы на отдельных location совпадают с заданием.
