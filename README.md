[![build](https://github.com/BurtsE/avito-test/actions/workflows/build.yml/badge.svg)](https://github.com/BurtsE/avito-test/actions/workflows/build.yml)
[![coverage](https://github.com/BurtsE/avito-test/actions/workflows/coverage.yml/badge.svg)](https://github.com/BurtsE/avito-test/actions/workflows/coverage.yml?percentage=23)
# Запуск

make docker-up 

либо 

docker-compose -f deploy/compose.yml up

# Список необходимых переменных окружения:

* USER_DB_USER="admin"
* USER_DB_PASSWORD=123
* HOUSE_DB_USER="admin"
* HOUSE_DB_PASSWORD=123
* HOUSE_DB="house_db"

# Дополнения к решению

* Обновление статуса квартиры по айди невозможно, если он не уникален, поэтому id и номер квартиры - разные поля
* Добавлен айди модератора у сущности квартиры
