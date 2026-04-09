# Отчёт

## Шаги выполнения
1. Прочитан [task.md](/home/savva/Desktop/devops/infra-course-labs/lab1/task.md) и подготовлен план выполнения.
2. Созданы каталоги `assets/screenshots` и `artifacts/logs`.
3. После получения `sudo` установлены зависимости:
   - `virtualbox`
   - `vagrant`
   - `golang-go`
4. Создан рабочий каталог [vagrant](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant) и подготовлены файлы [Vagrantfile](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant/Vagrantfile) и [main.go](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant/main.go).
5. Выполнена проверка конфигурации: `vagrant validate` завершился успешно.
6. Выполнен `vagrant up`:
   - box `ubuntu/jammy64` был успешно скачан и добавлен;
   - VirtualBox импортировал базовую машину и дошёл до шага `Booting VM`;
   - запуск ВМ завершился ошибкой гипервизора `AMD-V is not available (VERR_SVM_NO_SVM)`.
7. Выполнен `vagrant destroy -f`, но VirtualBox вернул ошибку доступа к объекту уже после неуспешного старта; состояние Vagrant осталось `poweroff`.

Краткий план:
- проверить среду;
- подготовить конфигурацию Vagrant;
- установить зависимости;
- запустить ВМ и зафиксировать фактический результат;
- задокументировать блокировку и выполнимую часть.

## Команды
- `mkdir -p lab1/assets/screenshots lab1/artifacts/logs`
- `sudo -n apt-get install -y golang-go virtualbox`
- `curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo -n gpg --dearmor --batch --yes -o /usr/share/keyrings/hashicorp-archive-keyring.gpg`
- `sudo -n apt-get install -y vagrant`
- `vagrant -v`
- `VBoxManage -v`
- `go version`
- `mkdir -p lab1/vagrant`
- `vagrant validate`
- `vagrant up`
- `vagrant status`
- `vagrant box list`
- `vagrant destroy -f`

## Вывод
- Хостовая ОС: Ubuntu 24.04.4 LTS.
- `vagrant`, `VirtualBox` и `go` успешно установлены и доступны в системе.
- Конфигурация [Vagrantfile](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant/Vagrantfile) валидна.
- Box `ubuntu/jammy64` успешно добавлен в локальную базу Vagrant.
- Основной блокер теперь подтверждён точно: текущая машина сама работает как гостевая VM под гипервизором `KVM`, при этом nested virtualization для AMD не проброшена внутрь гостя, поэтому VirtualBox внутри неё не может стартовать новую VM.
- Точная ошибка запуска:
  - `VBoxManage: error: AMD-V is not available (VERR_SVM_NO_SVM)`
- Из-за этого шаги `vagrant ssh`, `vagrant reload`, provision внутри гостевой ВМ и проверка `curl localhost:8080` не могут быть выполнены в текущей среде.
- Подготовлена конфигурация лабораторной:
  - [Vagrantfile](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant/Vagrantfile) с пробросом порта `80 -> 8080`;
  - shell provisioner устанавливает `golang-go` и регистрирует systemd-сервис `hello-vagrant.service`;
  - [main.go](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant/main.go) возвращает `Hello!` на HTTP-запрос.

Актуализация устаревших инструкций:
- В исходном задании используется `ubuntu/lunar64`.
- `lunar` является EOL-релизом Ubuntu, поэтому в подготовленной конфигурации заменён на `ubuntu/jammy64` как LTS-эквивалент, более подходящий для воспроизводимого запуска.
- Также запуск приложения через фоновой `go run ... &` заменён на systemd unit, чтобы сервис переживал переподключения shell provisioner и корректно рестартовал.

## Скриншоты
- Графическая среда и `scrot` не использовались.
- Вместо скриншотов сохранены текстовые артефакты:
  - [session_lab1_vagrant_up.typescript](/home/savva/Desktop/devops/infra-course-labs/lab1/artifacts/logs/session_lab1_vagrant_up.typescript)
  - [session_lab1_vagrant_destroy.typescript](/home/savva/Desktop/devops/infra-course-labs/lab1/artifacts/logs/session_lab1_vagrant_destroy.typescript)
