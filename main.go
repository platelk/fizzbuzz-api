package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/platelk/fizzbuzz-api/service"
)

const AppName = "FizzBuzz"
const DefaultPort = 8080

type AppParameters struct {
	Port uint16
}

func parseParams() *AppParameters {
	params := &AppParameters{}
	p := flag.String("port", os.Getenv("PORT"), "HTTP Port")

	flag.Parse()

	if p == nil || len(*p) == 0 {
		params.Port = DefaultPort
	} else {
		v, err := strconv.Atoi(*p)
		if err != nil {
			log.Fatalf("Invalid format for parameter [port].")
		}
		params.Port = uint16(v)
	}

	return params
}

func main() {
	log.Printf("Running %s httpService...", AppName)
	params := parseParams()
	log.Println("Create service...")
	httpService := service.CreateFizzBuzzService(params.Port)
	log.Println("Launch service...")
	httpService.Launch()
}
