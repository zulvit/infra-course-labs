# Отчёт

## Шаги выполнения
1. Подготовлены каталоги `assets/screenshots`, `artifacts/logs`, `work`.
2. Проверено наличие Docker: `docker` и Docker daemon отсутствуют в системе.
3. Для обхода ограничения без root установлен user-space стек в домашний каталог:
   - `pip` через `get-pip.py`;
   - `udocker` как замена Docker daemonless/runtime;
   - `ansible` для следующих лабораторных.
4. В [work/test.py](/home/savva/Desktop/devops/infra-course-labs/lab3/work/test.py) подготовлен HTTP-пример, аналогичный заданию, и в [work/run.sh](/home/savva/Desktop/devops/infra-course-labs/lab3/work/run.sh) собран сценарий воспроизведения.
5. Через `script` выполнен сценарий `lab3`:
   - `udocker pull python:3.11.5-alpine`
   - `udocker images`
   - `udocker run python:3.11.5-alpine python --version`
   - `udocker create --name=pyver ...`
   - запуск HTTP-сервера из контейнера с bind-mount каталога `work`
   - проверка ответа `wget -qO- http://127.0.0.1:8888`
   - `udocker inspect pyver`
   - `udocker manifest inspect python:3.11.5-alpine`
   - `udocker save -o python-3.11.5-alpine.tar ...`
6. Ошибки по ходу выполнения исправлялись точечно:
   - первый прогон оборвался на `udocker inspect -p image`, команда заменена на `udocker inspect pyver` и `udocker manifest inspect`.
   - второй прогон попал в конфликт уже занятого порта `8888`; в сценарий добавлен `pkill -f /vagrant/test.py`.
   - финальный шаг `save` был вынесен в отдельный лог, потому что tar-файл уже существовал после предыдущего успешного сохранения.

## Команды
- `docker --version`
- `docker info`
- `python3 /home/savva/.local/tmp/get-pip.py --user --break-system-packages`
- `pip3 install --user --break-system-packages ansible`
- `pip3 install --user --break-system-packages udocker`
- `script -qec 'bash lab3/work/run.sh' lab3/artifacts/logs/session_lab3.typescript`
- `script -qec '... udocker save ...' lab3/artifacts/logs/step04_udocker_save.typescript`

## Вывод
- Нативный Docker в системе отсутствует: `docker: command not found`.
- Для выполнения практики использован `udocker`, так как он не требует root и системного Docker daemon.
- Подтверждены базовые сценарии жизненного цикла контейнера:
  - загрузка образа `python:3.11.5-alpine`;
  - запуск контейнера и выполнение `python --version` с результатом `Python 3.11.5`;
  - создание именованных контейнеров `pyver` и `test-http`;
  - bind-mount каталога [work](/home/savva/Desktop/devops/infra-course-labs/lab3/work) в `/vagrant`;
  - запуск HTTP-сервера в контейнере и ответ с хоста `Hello from lab3 via udocker.`;
  - просмотр метаданных контейнера и manifest list образа;
  - сохранение образа в tar-архив размером `19M`.

Ограничения и актуализация:
- Исходное задание ориентировано на полноценный Docker daemon и команды `docker ...`.
- В текущей среде без root это недоступно, поэтому использован `udocker` как функционально близкий CLI для pull/run/create/inspect/save.
- Из-за этого разделы `docker exec`, `docker logs`, `docker top`, `docker stats`, локальный registry и TCP-доступ к daemon в точности по курсу не воспроизведены в `lab3`.
- Эти отличия являются осознанной адаптацией под среду, а не пропуском без попытки выполнения.

## Скриншоты
- Вместо графических скриншотов сохранены текстовые логи:
  - [session_lab3.typescript](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/session_lab3.typescript)
  - [step02_udocker_pull_run.typescript](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/step02_udocker_pull_run.typescript)
  - [step03_http_response.out](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/step03_http_response.out)
  - [step04_udocker_save.typescript](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/step04_udocker_save.typescript)
  - [python-3.11.5-alpine.tar](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/python-3.11.5-alpine.tar)
  - [step01_docker_version.err](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/step01_docker_version.err)
  - [step01_docker_info.err](/home/savva/Desktop/devops/infra-course-labs/lab3/artifacts/logs/step01_docker_info.err)
