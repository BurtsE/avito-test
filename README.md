[![coverage](https://github.com/BurtsE/avito-test/actions/workflows/coverage.yml/badge.svg)](https://github.com/BurtsE/avito-test/actions/workflows/coverage.yml)
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
* Для обеспечения большего RPS в сервис добавлен кэш. Фильтрация запроса по статусу модерации происходит на стадии запроса к БД, данные пользователя содержатся в контексте. Также для более быстрого ответа на запрос в таблице flats создан хэш индекс по аттрибуту house_id. 
