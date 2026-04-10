# Отчёт по лабораторной работе 4

## Данные по работе

- Студент: `Купленик С. С.`
- Группа: `УВП-171`
- Преподаватель: `Старший преподаватель кафедры ЦТУТП Заманов Е. А.`
- Дисциплина: `Инфраструктурное программное обеспечение`

## Цель работы

Цель работы состояла в развёртывании физической репликации PostgreSQL по схеме master-slave и проверке поведения реплики в основных эксплуатационных сценариях. Требовалось настроить пользователя репликации, physical replication slot, выполнить `pg_basebackup`, проверить потоковую репликацию, read-only режим реплики, отработать сценарии отставания и промоута.

## Ход работы

В лабораторной работе были подготовлены два контейнера PostgreSQL: `postgres_master` и `postgres_slave`. Для мастер-узла заданы параметры `wal_level`, `max_wal_senders`, `max_replication_slots`, а также правила доступа в `pg_hba.conf`. На реплике настроен автоматический `pg_basebackup` при старте контейнера и последующее подключение к мастеру.

После запуска стенда мастер увидел подключённую реплику. Ответ на запрос к `pg_stat_replication` имел вид:

```text
application_name | state     | replication_lag | sync_state
walreceiver      | streaming | 0 bytes         | async
```

Далее была проверена доставка данных: вставка строки на мастере появилась на реплике, а попытка выполнить `INSERT` непосредственно на slave завершилась ошибкой `cannot execute INSERT in a read-only transaction`, что подтвердило корректный режим горячего резерва.

В рамках самостоятельных заданий была выполнена массовая вставка 10 000 строк, смоделировано отключение реплики и проверено удержание WAL через replication slot, а затем выполнен `pg_promote()` с превращением бывшей реплики в новый мастер.

## Основные команды

Проверка состояния репликации:

```powershell
docker exec -e PGPASSWORD=shop_password_2026 postgres_master psql -U shop_admin -d shop_db -c "SELECT application_name, state, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), replay_lsn)) AS replication_lag, sync_state FROM pg_stat_replication;"
```

Проверка read-only режима на реплике:

```powershell
docker exec -e PGPASSWORD=shop_password_2026 postgres_slave psql -U shop_admin -d shop_db -c "INSERT INTO shop.manufacturers(manufacturer_name) VALUES ('Fail Vendor');"
```

Промоут реплики:

```powershell
docker exec -e PGPASSWORD=shop_password_2026 postgres_slave psql -U shop_admin -d shop_db -c "SELECT pg_promote(); SELECT pg_is_in_recovery();"
```

## Результаты

- потоковая репликация успешно запущена;
- мастер фиксирует реплику в состоянии `streaming`;
- реплика корректно работает в режиме только для чтения;
- массовая вставка реплицируется с минимальным лагом в локальной сети Docker;
- replication slot удерживает WAL во время простоя реплики;
- после `pg_promote()` реплика становится новым мастером и принимает запись.

## Итог

Лабораторная работа выполнена полностью. Настроена и проверена физическая репликация PostgreSQL, подтверждена доставка изменений на реплику, поведение системы при временном отключении standby и корректность промоута резервного узла в новый мастер.


