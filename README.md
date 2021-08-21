# Cinema Backend

```sh
go get github.com/labstack/echo/v4

go get go.mongodb.org/mongo-driver
go get github.com/gobeam/mongo-go-pagination

go get github.com/dgrijalva/jwt-go
go get golang.org/x/crypto
go get golang.org/x/sys

go get github.com/swaggo/echo-swagger
go get github.com/swaggo/swag
go get github.com/swaggo/swag/cmd/swag

go get github.com/go-playground/universal-translator
go get github.com/go-playground/validator
```
* Install dependencies:

```sh
go mod download
```

* Generate Swagger doc:

```sh
swag init
```

## Run Docker Compose

* Run

```sh
docker-compose up -d 
```

* Down

```sh
docker-compose down 
```