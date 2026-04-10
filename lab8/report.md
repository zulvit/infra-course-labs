# Отчёт по лабораторной работе №8

## Титульные данные

- Студент: Купленик С. С.
- Группа: УВП-171
- Преподаватель: ст. преп. каф. ЦТУТП Заманов Е. А.

## Цель

Оформить установку nginx как роль Ansible, применить её к узлам, переопределить файлы и переменные вторым плейбуком, затем поставить коллекцию `nginxinc.nginx_core` и сравнить сценарии «из коробки» и со своим шаблоном страницы.

## Ход работы

Роль создана через `ansible-galaxy role init`, доработаны `roles/nginx/tasks`, `handlers`, `defaults`, `files/default.conf`, `templates/index.html.j2`. Плейбуки: `playbook.yaml`, `playbook_override.yaml`, для коллекции — `playbook_collection.yaml`, `playbook_collection_demo.yaml`.

Каталог практики копировался на bastion (стенд тот же, что после лаб. 7). Запуски с bastion, инвентарь в виде списка хостов:

```bash
ansible-galaxy collection install nginxinc.nginx_core
ansible-playbook -i node1.local,node2.local playbook.yaml
ansible-playbook -i node1.local,node2.local playbook_override.yaml
ansible-playbook -i node1.local,node2.local playbook_collection.yaml
ansible-playbook -i node1.local,node2.local playbook_collection_demo.yaml
```

Проверки `curl` и просмотр файлов на node1 — в логах `artifacts/logs/step07_*.log` и далее по номерам.

## Результаты (фрагменты вывода)

После второго плейбука (другой конфиг и шаблон, порт 8080) тело ответа:

```text
### overrided hello from node1.local !
```

(В ответе видна опечатка в слове *overridden* — так отдал сервер, в конфиге шаблона то же.)

После демо-плейбука коллекции со своим шаблоном главная страница:

```text
hello from node1.local
```

## Замечания

На jammy из коробки ставится ansible 2.10.x; коллекция при установке может ругаться на версию, но плейбуки отработали. Если в YAML после `role init` оказалось два маркера `---`, плейбук падает — нужно оставить один в начале файла.

## Вывод

Собственная роль накатывается на оба узла, переопределение переменных и handler перезапуска nginx работают. Коллекция ставится зависимостями целиком, стандартная страница и замена на свой шаблон проверены запросами HTTP.
