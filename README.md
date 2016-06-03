A Simple CSV Upload and Query App
=============================

## Specs

- [x] should upload csv files to server
- [x] server should be able to store the csv data in Redis
- [x] data in Redis should be stored progressively
- [x] should expose API for querying the DB
- [x] should unit tests

## Usage

Installation
-----------
```sh
$ go install
```

Running Server
-----------
```sh
$ go run main.go
```

Running Client
-----------
```sh
$ go run upload.go
```

APIs Exposed
-----------
```json
Upload CSV - 
POST - localhost:9000/upload
```

``` QUERIES
GET - localhost:9000/query?timestamp=146293783&type=ObjectB
GET - localhost:9000/query?timestamp=146293783&type=*
```

Testing
-----------
```sh
$ go test test/redis_test.go
```

## MVC Structure

```json 
Project Struct
├── LICENSE
├── README.md
├── conf
│   ├── app.conf
│   └── config.go
├── controllers
│   ├── query.go
│   ├── sayHello.go
│   ├── static.go
│   └── uploadCSV.go
├── core
│   └── router.go
├── main.go
├── middlewares
│   └── logger.go
├── models
│   ├── dao.go
│   ├── dbutils.go
│   └── event.go
├── routers
│   └── router.go
├── utils
│   ├── base64util.go
│   ├── md5util.go
│   └── stringutil.go
└── test
    └── redis_test.go
 ```