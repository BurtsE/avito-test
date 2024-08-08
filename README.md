[![CI](https://github.com/BurtsE/avito-test/actions/workflows/ci.yml/badge.svg)](https://github.com/BurtsE/avito-test/actions/workflows/ci.yml)

# Запуск

make docker-up 
либо 
docker-compose -f deploy/compose.yml up

# Список необходимых переменных окружения:

USER_DB_USER="admin"\n
USER_DB_PASSWORD=123\n
HOUSE_DB_USER="admin"\n
HOUSE_DB_PASSWORD=123\n
HOUSE_DB="house_db"\n
