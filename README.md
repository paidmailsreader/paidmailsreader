# PaidMailsReader key server

## Установка

Скачать последнюю версию [Node.js](https://nodejs.org/ru/)

Скачать последнюю версию [Git](https://git-scm.com/)

Открыть консоль Windows и выполнить команды

```bash
git clone https://github.com/kas-cor/pmr-key-server.git
cd pmr-key-server
npm install
```

## Настройка хоста

1. Запустите Блокнот(notepad) от имени Администратора
2. Файл->Открыть..., в поле "Имя файла:" вставьте `C:\Windows\System32\drivers\etc\hosts`, где `C` - системный диск
3. В конец файла добавьте строку `127.0.0.1 pmr.kas-cor.ru`
4. Закройте файл и подтвердите сохранение

## Запуск сервера

Открыть консоль Windows, перейти в папку с сервером и выполнить команду

```bash
npm start
```
