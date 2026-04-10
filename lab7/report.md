# Отчёт

## Данные по работе

- Студент: `Купленик С. С.`
- Группа: `УВП-171`
- Преподаватель: `Старший преподаватель кафедры ЦТУТП Заманов Е. А.`

## РџРѕРґСЂРѕР±РЅС‹Рµ РѕС‚РІРµС‚С‹ РєРѕРјР°РЅРґ

РљРѕРјР°РЅРґР°:

```powershell
ansible -i /vagrant/hosts -m ping all
```

РЎРјС‹СЃР» РїРѕРґС‚РІРµСЂР¶РґС‘РЅРЅРѕРіРѕ РѕС‚РІРµС‚Р°:

```text
bastion.local | SUCCESS => {"ping": "pong"}
node1.local   | SUCCESS => {"ping": "pong"}
node2.local   | SUCCESS => {"ping": "pong"}
```

РљРѕРјР°РЅРґР°:

```powershell
curl node1.local/node1only
```

РћС‚РІРµС‚:

```text
node1 only
```





