# FizzBuzz HTTP Service

FizzBuzz Service provide a http REST API to resolve FizzBuzz commonly know technical test.

## Prerequisite

To compile this project you need:
 * `go` installed
 
## Setup

### 1. Install deps

If you have [glide](https://glide.sh/) instal you can run

> glide i

Otherwise, you can run

> go get ./...

### 2. Build

* For normal use

> go build

* For Docker image

> GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo


## Usage

To run, simply execute it:

> ./fizzbuzz-api

To display it's usage, simply run:

> ./fizzbuzz-api -h

    Usage of ./fizzbuzz-api:
        -port string
    	         HTTP Port

## API

`fizzbuzz-api` have several routes available

### /fizzbuzz/v1/version route

![](https://img.shields.io/badge/Request-GET-green.svg?style=flat)

```Curl
<!--Curl-->
http://<host>:<port>/fizzbuzz/v1/version
```


- **200 Ok**
```json
HTTP/1.1 200 OK
{"statusCode":0,"version":"0.0.1"}
```

### /fizzbuzz/v1/fizzbuzz

![](https://img.shields.io/badge/Request-GET-green.svg?style=flat)

```Curl
<!--Curl-->
http://<host>:<port>/fizzbuzz/v1/fizzbuzz?to=<to>&mul1=<mul1>&mul2=<mul2>&word1=<word1>&word2=<word2>
```
| Field | Description
| ----- | -----------
| to | limit of the fizzbuzz
| mul1 | first multiple which will be replace by <word1>
| mul2 | second multiple which will be replace by <word2>
| word1 | word to replace number multiple of <mul1>
| word2 | word to replace number multiple of <mul2>


- **200 OK**
```json
HTTP/1.1
{
    "Response": [
        "1",
        "2",
        "fizz",
        "4",
        "buzz",
        "fizz",
        "7",
        "8",
        "fizz",
        "buzz",
        "11",
        "fizz",
        "13",
        "14",
        "fizzbuzz"
    ],
    "statusCode": 0
}
```

- **406 Invalid agurment**
```json
HTTP/1.1 406
{
    "message": "One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)",
    "status": "invalid argument",
    "statusCode": 1
}
```
