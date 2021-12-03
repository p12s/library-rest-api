<div align="center">
<article style="display: flex; flex-direction: column; align-items: center; justify-content: center;">
  <h1 style="width: 100%; text-align: center;">Library REST-API</h1>
  <p>Simple REST-API server. Task description is <a href="task.md">here</a></p>
</article>

<div align="center">

[![LICENSE][license-badge]][license-url]
  
[license-badge]: https://img.shields.io/npm/l/@douyinfe/semi-ui
[license-url]: https://github.com/p12s/library-rest-api/blob/master/LICENSE
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

# What is going on here  
  
1. What does the library do:  
- can register/authorize a user  
- can create/update/delete the book  
  
2. What does the asuncronus logger (gRPC-client + gRPC-server) do:  
- can send any handler-action to logger-service  
  
3. What does the asuncronus logger-2 (RabbitMQ-producer + RabbitMQ-consumer) do:   
- can send any handler-action to logger-service  
  
# Logging services
- **logger** - logging with gRPC, saving into MongoDB  
- **logger-2** - logging with RabbitMQ, simple printing into stdout  

# 📌 Documentation
* [OpenAPI](docs/README.md)
