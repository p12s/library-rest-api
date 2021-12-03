FROM golang:1.17.2-buster AS build

RUN go version
ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o /app ./cmd/main.go

FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /app /app
COPY ./.env /.env
COPY ./wait-for-postgres.sh ./

# install psql and make wait-for-postgres.sh executable
RUN apk add --no-cache libc6-compat postgresql-client && chmod +x wait-for-postgres.sh app

CMD ["./app"]