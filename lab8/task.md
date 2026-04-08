# Ansible Roles and Collections

**Источник курса:** [08_practice_ansible_roles.html](https://yudolevich.github.io/infra-course/08_practice_ansible_roles.html)

## Задание

Практическое занятие: роли (roles) и коллекции (collections) в Ansible.

### Vagrant

Тот же стек, что в практике Ansible: `bastion`, `node1`, `node2`, ключ `key`, `install` ключа на bastion, `ANSIBLE_HOST_KEY_CHECKING=False` (см. лабораторную 7).

### Inventory через запятую

Пример ad-hoc:

```bash
ansible -m ping -i node1.local,node2.local all
```

### Roles

#### Init

Создание структуры роли: `ansible-galaxy role init --init-path roles nginx`. Каталоги: `defaults`, `files`, `handlers`, `meta`, `tasks`, `templates`, `tests`, `vars`, `README.md`.

#### Tasks

`roles/nginx/tasks/main.yml` — установка пакета `nginx`.

Плейбук с `roles: - name: nginx`. Запуск: `ansible-playbook -i node1.local,node2.local playbook.yaml`.

#### Defaults

`roles/nginx/defaults/main.yml`: переменные `nginx_html_dir`, `nginx_config_file`, `nginx_template_file`.

#### Files/Templates

`roles/nginx/files/default` — конфиг nginx; `roles/nginx/templates/index.html.j2` — шаблон страницы.

Задачи в `tasks/main.yml`: package, copy конфига, template `index.html`.

Переопределение через файлы рядом с плейбуком: `templates/overrided.html.j2`, `files/overrided.conf` и параметры роли в плейбуке (`nginx_template_file`, `nginx_config_file`). В курсе отмечается, что после смены конфига nginx может не подхватить порт 8080 без перезапуска.

#### Handlers

`roles/nginx/handlers/main.yml` — handler `nginx restart` (systemd). В задаче копирования конфига — `notify: nginx restart`. Обновить `overrided.conf` (добавить `listen 8080`), снова запустить плейбук, проверить `curl node1.local:8080`.

### Collections

Перед разделом Collections в курсе рекомендуется `vagrant destroy -f` и `vagrant up`; примеры коллекций — в пустой директории с нуля.

#### Install

`ansible-galaxy collection install nginxinc.nginx_core` (в логе курса также подтягиваются зависимости: `ansible.posix`, `community.crypto`, `community.general`).

#### Usage

В плейбуке: `collections: [nginxinc.nginx_core]`, роли `nginxinc.nginx_core.nginx` и `nginx_config` с переменной `nginx_config_http_template` (полный YAML — как в курсе). Запуск плейбука, проверка `curl` с `grep title`.

Дополнительно: шаблон `templates/index.html.j2` рядом с плейбуком и переменные `nginx_config_html_demo_template_enable`, `nginx_config_html_demo_template` для выкладки `index.html`.

## Требования

- Инфраструктура из **лабораторной 7** (или пересозданные ВМ по курсу).
- Доступ к **Ansible Galaxy** для установки коллекции `nginxinc.nginx_core`.

## Шаги выполнения (предложенные)

1. Создать роль `nginx`, подключить в плейбук, расширить defaults/files/templates/tasks/handlers по тексту курса.
2. Проверить переопределение переменных и перезапуск nginx через handler.
3. Установить коллекцию `nginxinc.nginx_core`, применить плейбук с ролями из коллекции, затем вариант с демо-шаблоном HTML.

## Ожидаемый результат

- Собственная роль настраивает nginx и отдаёт персонализированную страницу на node1/node2.
- После правок конфига и handler доступен порт 8080 (как в финальном примере курса).
- Коллекция `nginxinc.nginx_core` разворачивает nginx и конфиг; `curl` показывает ожидаемые заголовки/тело.

## Команды (выделено для CLI-агента)

```bash
ansible-galaxy role init --init-path roles nginx
ls -1 roles/nginx/
```

```bash
ansible-playbook -i node1.local,node2.local playbook.yaml
curl -s node1.local | grep title
curl -s node1.local
curl node1.local:8080
```

```bash
vagrant destroy -f
vagrant up
```

```bash
ansible-galaxy collection install nginxinc.nginx_core
ansible-playbook -i node1.local,node2.local playbook.yaml
curl -s node1.local | grep title
curl node1.local
```

Полные **`playbook.yaml`** (все варианты), файлы роли **`roles/nginx/**`**, **`overrided.*`**, плейбук с **collections** — с [страницы курса](https://yudolevich.github.io/infra-course/08_practice_ansible_roles.html). В тексте Galaxy коллекция названа **nginxinc.nginx_core** (роли `nginx`, `nginx_config`).
