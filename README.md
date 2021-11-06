![License](https://img.shields.io/github/license/p12s/library-rest-api)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/library-rest-api?style=plastic)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/library-rest-api)](https://goreportcard.com/report/github.com/p12s/library-rest-api)

# Library REST-API
Task description is [here](task.md)

## How to run
```
git clone https://github.com/p12s/library-rest-api.git && cd library-rest-api

mv .env.example .env

docker-compose up -d

open in browser: http://localhost/swagger/index.html

docker-compose down
```
There is a bug with migrations - they are not performed automatically as they should.  
You have to run the command manually.
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

## OpenAPI
[Documentation](docs/README.md)

