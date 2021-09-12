# PaidMailsReader key server

## Установка

Скачать и установить последнии версии [Node.js](https://nodejs.org/ru/) и [Git](https://git-scm.com/)

Открыть консоль Windows и выполнить команды

```bash
git clone https://github.com/kas-cor/pmr-key-server.git
cd pmr-key-server
npm install
```

## Настройка хоста

1. Запустите Блокнот(notepad.exe) от имени Администратора
2. Файл->Открыть..., в поле "Имя файла:" вставьте `C:\Windows\System32\drivers\etc\hosts`, где `C` - буква системного диск
3. В конец файла добавьте строку `127.0.0.1 pmr.kas-cor.ru`
4. Закройте Блокнот(notepad.exe) и подтвердите сохранение файла `hosts`

## Запуск сервера

Открыть консоль Windows, перейти в папку с сервером и выполнить команду

```bash
npm start
```

Для проверки работы сервера, перейдите по [ссылке](http://pmr.kas-cor.ru)

Если все сделано верно, вы увидите `Key server PMR is running...`

Для работы программы введите ключ `F1Z6S3L9Z4K9J3W9Z6K7D9B1A5G8`
