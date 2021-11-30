<div align="center">
<article style="display: flex; flex-direction: column; align-items: center; justify-content: center;">
  <h1 style="width: 100%; text-align: center;">Library REST-API</h1>
  <p>Simple REST-API server. Task description is <a href="task.md">here</a></p>
</article>

<div align="center">

[![GO][go-badge]][go-url] [![REPORT][report-badge]][report-url] [![LICENSE][license-badge]][license-url]
  
[license-badge]: https://img.shields.io/npm/l/@douyinfe/semi-ui
[license-url]: https://github.com/p12s/library-rest-api/blob/master/LICENSE
[go-badge]: https://img.shields.io/github/go-mod/go-version/p12s/library-rest-api?style=plastic
[go-url]: https://github.com/p12s/library-rest-api/blob/master/go.mod
[report-badge]: https://goreportcard.com/badge/github.com/p12s/library-rest-api
[report-url]: https://goreportcard.com/report/github.com/p12s/library-rest-api

</div>
</div>

# 🔥 Run
Download:
```sh
git clone https://github.com/p12s/library-rest-api.git && cd library-rest-api
mv library/.env.example library/.env
mv logger/.env.example logger/.env
mv logger-2/.env.example logger-2/.env

docker-compose up --build
```
Open in browser: http://localhost/swagger/index.html  
Stop:  
```sh
docker-compose down
```

There is a bug with migrations - they are not performed automatically as they should.  
You have to run the command manually.
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

# Logging
- service **logger** - logging with gRPC, saving into MongoDB  
- service 2 **logger-2** - logging with RabbitMQ, output into output  

# 📌 Documentation
* [OpenAPI](docs/README.md)
