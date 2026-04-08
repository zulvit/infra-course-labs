# Ansible

**Источник курса:** [07_practice_ansible.html](https://yudolevich.github.io/infra-course/07_practice_ansible.html)

## Задание

В данном практическом занятии рассматривается базовое использование Ansible.

### Vagrant

`Vagrantfile` с тремя машинами: `bastion` (ansible, libnss-mdns), `node1`, `node2` (libnss-mdns), частная сеть dhcp.

Управление с **bastion**. SSH-ключ: путь из `vagrant ssh-config bastion | grep IdentityFile`; скопировать в каталог проекта как `key`. На bastion:

```bash
install -m 600 -o vagrant /vagrant/key /home/vagrant/.ssh/id_rsa
export ANSIBLE_HOST_KEY_CHECKING=False
```

### Inventory

Файл `hosts` с группами `[bastion]`, `[nodes]` и хостами `*.local`.

### Basic Usage

Ad-hoc: `ansible -i hosts -m ping all`, `ansible -i hosts -m shell -a 'uname -a' nodes`, ограничение `-l`.

**Примечание:** `ansible-doc`, `ansible-doc -l` для справки по модулям.

### Playbook

#### Tasks

Плейбук: установка `nginx` на `nodes` модулем `ansible.builtin.package`, `become: true`. Запуск: `ansible-playbook -i hosts playbook.yaml`. Проверка: `curl -s node1.local | grep title`.

`gather_facts: false` — пропуск сбора фактов; повторный запуск — задачи в состоянии `ok` если изменений нет.

#### Handlers

Копирование конфигурации nginx (`copy` из `files/default` в `/etc/nginx/sites-enabled/default`), `notify: nginx restart`, handler с `systemd` для рестарта nginx.

#### Jinja2

Шаблон `templates/index.html.j2` с `hello from {{ ansible_host }}`, задача `template`, переменная `html_dir` в `vars`.

#### Conditionals

Задача с `when: ansible_host == "node1.local"` — файл только на node1.

#### Loops

Задача с `loop: [test1, test2, test3]` и `item` в `dest`.

## Требования

- Три ВМ по `Vagrantfile` курса; DNS-имена `*.local` (mdns).
- На **bastion**: установленный **ansible**.
- Ключ **`key`** в каталоге с `Vagrantfile`, размещённый в `~vagrant/.ssh/id_rsa` на bastion.

## Шаги выполнения (предложенные)

1. Поднять ВМ, настроить ключ и `ANSIBLE_HOST_KEY_CHECKING=False`.
2. Создать `hosts`, выполнить `ping` и ad-hoc команды.
3. Написать и запустить плейбук установки nginx; добавить handlers, шаблон, условия, циклы по мере прохождения разделов курса.

## Ожидаемый результат

- `ansible -i hosts -m ping all` — SUCCESS для всех хостов.
- После плейбука nginx отдаёт страницы; персонализированный `index.html` с именем хоста; условные и циклические задачи ведут себя как в примерах курса.

## Команды (выделено для CLI-агента)

```bash
vagrant ssh-config bastion | grep IdentityFile
```

```bash
install -m 600 -o vagrant /vagrant/key /home/vagrant/.ssh/id_rsa
export ANSIBLE_HOST_KEY_CHECKING=False
```

```bash
ansible -i hosts -m ping all
ansible -i hosts -m shell -a 'uname -a' nodes
ansible -i hosts -m shell -a 'hostname' nodes -l node1.local
ansible-playbook -i hosts playbook.yaml
curl -s node1.local | grep title
curl node1.local
curl node1.local/node1only
curl node1.local/test1
```

Содержимое **`Vagrantfile`**, **`hosts`**, **`playbook.yaml`** (все версии), **`files/default`**, **`templates/index.html.j2`** — с [страницы курса](https://yudolevich.github.io/infra-course/07_practice_ansible.html).
