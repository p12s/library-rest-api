# Ручной запуск контейнера с БД
### При запуске сервиса в docker-compose это делать не нужно
  
Запуск контейнера с postgresql (для простоты MVP база в нем):
```
docker pull postgres
docker run --name=library-rest-api -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
docker exec -it PID /bin/bash
```
Запускаем миграции:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up # вместо up -> down, если что-то пошло не так
```
В контейнере можно убедиться, что БД с таблицами созданы:
```
docker ps
docker exec -it PID /bin/bash

psql -U postgres
\d
select * from users;
```

Запуск приложения:
```
go run ./cmd/main.go
```

Документация методов:  
http://localhost:80/swagger/index.html
