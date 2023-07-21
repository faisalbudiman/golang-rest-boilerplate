# golang-rest-boilerplate
Restfull API in Golang Echo Framework. This project already implemented life-cycle programming and [clean code](http://cleancoder.com/).


This project features are:
- gorm (already used with postgres/mysql)
- elasticsearch (default disabled)
- logging with logrus
- Authentication with JWT
- example Rest API (CRUD) with and without authentication (token)
- migrations
- request validation
- swagger docs
- environment configuration
- redis
- docker development environment


## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run

Local environment
``` bash
# build 
$ go build

# run in development 
$ ENV=DEV go run main.go
$ ENV=DEV ./filego

# run in staging 
$ ENV=STAGING go run main.go
$ ENV=STAGING ./filego

# run in production 
$ ENV=PROD go run main.go
$ ENV=PROD ./filego
```

Docker environment

in docker environment, migration will get error because the table of 'sample' used UUID v4, please run `CREATE EXTENSION "uuid-ossp";` for postgres.
``` bash
# build 
$ docker build -t golang-rest-boilerplate:latest .

# config
sudo sysctl -w vm.max_map_count=262144 # it is required for elasticsearch config

# run
$ docker compose up
```

## Documentation

Install environment
``` bash
# get swagger package 
$ go get github.com/swaggo/swag

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{url}}/swagger/index.html

the project will be run in:
- PORT : 3030
- PATH : /

# Libraries
- Echo - https://github.com/labstack/echo/v4
- ORM (gorm) - https://gorm.io/gorm
- JWT - https://github.com/dgrijalva/jwt-go
- Swagger - https://github.com/swaggo/echo-swagger
- Logger - https://github.com/sirupsen/logrus

# Author
Faisal Budiman