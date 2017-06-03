package service

import (
	"fmt"
	"log"
	"net/http"
	"github.com/platelk/fizzbuzz-api/core"
)

// HttpService define basic possible interaction for all HttpService
type HttpService interface {
	Launch()
}

// FizzBuzzService is a HttpService which
type FizzBuzzService struct {
	httpClient *http.ServeMux
	version    string
}

func CreateFizzBuzzService() HttpService {
	return &FizzBuzzService{
		version: "0.0.1",
	}
}

// routes will setup all the available route
func (service *FizzBuzzService) routes() http.Handler {
	service.httpClient = http.NewServeMux()

	service.httpClient.HandleFunc("/fizzbuzz/v1/version", service.VersionRoute)
	service.httpClient.HandleFunc("/fizzbuzz/v1/fizzbuzz", service.FizzBuzzRoute)

	return service.httpClient
}

func (service *FizzBuzzService) Launch() {
	log.Println("Service http service...")
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), applyCORS(service.routes()))
}

func (service *FizzBuzzService) VersionRoute(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data, err := MessageToJson(CreateVersionMessage(service.version))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			data, err = MessageToJson(CreateErrorMessage("serialization error", "Error during json serialization."))
			resp.Write(data)
		}
		resp.Write(data)
	default:
		data, err := MessageToJson(CreateErrorMessage("wrong method", "/version route only accept GET method."))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			data, err = MessageToJson(CreateErrorMessage("serialization error", "Error during json serialization."))
			resp.Write(data)
		}
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write(data)
	}
}

func (service *FizzBuzzService) FizzBuzzRoute(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		respond, err := core.FizzBuzz(1, 30, 3, 5, "fizz", "buzz")
		if err != nil {
			resp.WriteHeader(http.StatusNotAcceptable)
			data, _ := MessageToJson(CreateErrorMessage("invalid argument", err.Error()))
			resp.Write(data)
		}
		data, err := MessageToJson(CreateFizzBuzzMessage(respond))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			data, err = MessageToJson(CreateErrorMessage("serialization error", "Error during json serialization."))
			resp.Write(data)
		}
		resp.Write(data)
	default:
		data, err := MessageToJson(CreateErrorMessage("wrong method", "/fizzbuzz route only accept GET method."))
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			data, err = MessageToJson(CreateErrorMessage("serialization error", "Error during json serialization."))
			resp.Write(data)
		}
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write(data)
	}
}
