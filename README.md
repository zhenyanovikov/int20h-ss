# Project Name

## Description


## Installation
1. Склонуйте репозиторій (`git clone`)
2. Встановіть [docker](https://docs.docker.com/get-docker/), [docker-compose](https://docs.docker.com/compose/install/)
та [migrate](https://github.com/golang-migrate/migrate) 
3. Скопіюйте конфігурацію `backend/config_example.json` як `backend/config.json`, та відредагуйте його відповідно до вашої конфігурації
4. Виконайте команду `docker-compose up` в кореневій папці проєкту
5. Виконайте команду `cd backend && DB_PASSWORD=<пароль до бд> DB_PORT=5432 make migrate-up` для проведення міграцій
6. Відкрийте `http://localhost:8080` в браузері
7. Насолоджуйтесь!
