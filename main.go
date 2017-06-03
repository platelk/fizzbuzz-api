package main

import (
	"log"
	"github.com/platelk/fizzbuzz-api/service"
)

const AppName = "FizzBuzz"

func main() {
	log.Printf("Running %s httpService...", AppName)
	httpService := service.CreateFizzBuzzService()
	httpService.Launch()
}
