git pull --rebase

docker-compose build backend
docker-compose build frontend
docker-compose -f docker-compose.yml up -d
