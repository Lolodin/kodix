# Readme
CRUD API
* `/product/{id}` - получаем json о продукте
* `/List` - массив с данными всех продуктов
* `/delete/{id}` - удаление продукта из бд
* `/put` - обновляем данные о продукте
* `/new` - добавляем новый продукт

При переходе на localhost:8080 будет простое React приложение где можно протестировать работу сервера

## Запуск

`$ git clone https://github.com/Lolodin/kodix`

`$ cd kodix`

`$ go build cmd/main.go `

`$ /main`

Откройте браузер по адресу http://localhost:8080/ 

## Запуск в Docker

`$ git clone https://github.com/Lolodin/kodix`

`$ cd kodix`

`$ docker build -t test .`

`$ docker run -it -p 8080:8080 test`

Откройте браузер по адресу http://localhost:8080/ 

