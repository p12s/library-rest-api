<div align="center">
  <article style="display: flex; flex-direction: column; align-items: center; justify-content: center;">
    <h1 style="width: 100%; text-align: center;">Library REST-API</h1>
    <p>
      Simple REST-API server. Task description is [here](task.md)
    </p>
  </article>
</div>

<div align="center">
  [![GO][go-badge]] [![REPORT][report-badge]][report-url] [![LICENSE][license-badge]][license-url]
  
  [license-badge]: https://img.shields.io/npm/l/@douyinfe/semi-ui
  [license-url]: https://github.com/p12s/library-rest-api/blob/master/LICENSE
  
  [go-badge]: https://img.shields.io/github/go-mod/go-version/p12s/library-rest-api?style=plastic
  
  [report-badge]: https://goreportcard.com/badge/github.com/p12s/library-rest-api
  [report-url]: https://goreportcard.com/report/github.com/p12s/library-rest-api
</div>

# 🔥 Run
Download:
```sh
git clone https://github.com/p12s/library-rest-api.git && cd library-rest-api
mv .env.example .env
docker-compose up -d
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

# 📌 Documentation
* [OpenAPI](docs/README.md)
