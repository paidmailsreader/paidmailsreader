# PaidMailsReader key server

## Настройка хоста

1. Запустите Блокнот(notepad.exe) от имени Администратора
2. Файл->Открыть..., в поле "Имя файла:" вставьте `C:\Windows\System32\drivers\etc\hosts`, где `C` - буква системного диска
3. В конец файла добавьте строку `127.0.0.1 pmr.kas-cor.ru`
4. Закройте Блокнот(notepad.exe) и подтвердите сохранение файла `hosts`

## Запуск сервера

Запустите `pmr-key-server.exe`

Для работы программы введите ключ `F1Z6S3L9Z4K9J3W9Z6K7D9B1A5G8`

## Сборка сервера

1. Скачать и установить последнюю версию [Golang](https://go.dev/dl/)
2. Скачать архив исходников [скачать](https://github.com/paidmailsreader/paidmailsreader.github.io/archive/refs/heads/master.zip)
3. Распаковать в отдельную директорию
4. Открыть консоль Windows и в директории выполнить команду

```bash
GOOS=windows GOARCH=amd64 go build -o build/pmr-key-server.exe main.go
```

Исполняемый файл-сервера ключей будет в папке `build`
