# Отчёт

## Шаги выполнения
1. Прочитан [task.md](/home/savva/Desktop/devops/infra-course-labs/lab2/task.md) и оценены требования к среде.
2. Проверен базовый Vagrant/VirtualBox стек на предыдущей лабораторной:
   - `vagrant` установлен;
   - `VirtualBox` установлен;
   - `go` установлен.
3. Выполнена фактическая проверка возможности запуска вложенной виртуальной машины через `vagrant up` в `lab1`.
4. Проверка завершилась ошибкой гипервизора `AMD-V is not available (VERR_SVM_NO_SVM)`, то есть nested virtualization в текущей гостевой системе не предоставлена.
5. Поскольку `lab2` требует уже не одну, а три виртуальные машины (`front`, `back`, `db`) и работу `private_network`, лабораторная заблокирована тем же системным ограничением ещё до этапа сборки боксов `front/back/db`.

## Команды
- `vagrant -v`
- `VBoxManage -v`
- `go version`
- `vagrant up` в [lab1/vagrant](/home/savva/Desktop/devops/infra-course-labs/lab1/vagrant) как проверка базовой работоспособности гипервизора
- `lscpu`
- `egrep -c '(vmx|svm)' /proc/cpuinfo`

## Вывод
- В текущей среде можно установить `Vagrant` и `VirtualBox`, но нельзя запустить гостевые ВМ внутри уже работающей ВМ, потому что не доступно аппаратное ускорение AMD-V.
- Для `lab2` это критично:
  - нельзя поднять базовую машину;
  - нельзя выполнить `vagrant package`;
  - нельзя собрать локальные боксы `front`, `back`, `db`;
  - нельзя поднять multi-node стенд и проверить `http://localhost:8888`.
- Следовательно, `lab2` в исходном формате честно невыполнима в текущей гостевой системе без включения nested virtualization на внешнем гипервизоре или переноса работы на физический хост/VM с проброшенным SVM/VT-x.

## Скриншоты
- Вместо графических скриншотов использованы текстовые артефакты из фактической проверки Vagrant/VirtualBox:
  - [session_lab1_vagrant_up.typescript](/home/savva/Desktop/devops/infra-course-labs/lab1/artifacts/logs/session_lab1_vagrant_up.typescript)
  - [session_lab1_vagrant_destroy.typescript](/home/savva/Desktop/devops/infra-course-labs/lab1/artifacts/logs/session_lab1_vagrant_destroy.typescript)
