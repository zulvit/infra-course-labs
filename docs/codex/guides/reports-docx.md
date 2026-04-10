# REPORTS_DOCX

После установки `otchet-compose` генерация всех восьми DOCX выполняется из корня проекта:

```powershell
1..8 | ForEach-Object { Set-Location "lab$_"; otchet-compose gen -f otchet-compose.yml; Set-Location .. }
```

Перед генерацией нужно установить утилиту:

```powershell
pip install -e ./otchet-compose
```

В титульных листах нужно заменить плейсхолдеры:

- `TODO_INSTITUTE`
- `TODO_DEPARTMENT`
- `TODO_DISCIPLINE`
- `TODO_STUDENT`
- `TODO_GROUP`
- `TODO_TEACHER`

| Лаба | YAML | DOCX |
|---|---|---|
| lab1 | `lab1/otchet-compose.yml` | `lab1/build/lab1-report.docx` |
| lab2 | `lab2/otchet-compose.yml` | `lab2/build/lab2-report.docx` |
| lab3 | `lab3/otchet-compose.yml` | `lab3/build/lab3-report.docx` |
| lab4 | `lab4/otchet-compose.yml` | `lab4/build/lab4-report.docx` |
| lab5 | `lab5/otchet-compose.yml` | `lab5/build/lab5-report.docx` |
| lab6 | `lab6/otchet-compose.yml` | `lab6/build/lab6-report.docx` |
| lab7 | `lab7/otchet-compose.yml` | `lab7/build/lab7-report.docx` |
| lab8 | `lab8/otchet-compose.yml` | `lab8/build/lab8-report.docx` |
