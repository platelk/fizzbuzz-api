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

> go build

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

### Version

- *GET* `/fizzbuzz/v1/version`