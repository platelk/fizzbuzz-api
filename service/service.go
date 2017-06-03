package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bluele/gcache"
	"github.com/platelk/fizzbuzz-api/core"
)

// FizzBuzzRespCacheSize represent the number of response that will be cached
const FizzBuzzRespCacheSize = 5000
// Version is a string representation of the service version
const Version = "0.0.1"

// HttpService define basic possible interaction for all HttpService
type HttpService interface {
	// Launch the service
	Launch()
}

// FizzBuzzService is a HttpService which
type FizzBuzzService struct {
	httpClient         *http.ServeMux
	version            string
	port               uint16
	fizzBuzzRouteCache gcache.Cache
}

// CreateFizzBuzzService initialize a new FizzBuzzService
func CreateFizzBuzzService(port uint16) HttpService {
	return &FizzBuzzService{
		version: Version,
		port: port,
		fizzBuzzRouteCache: gcache.
			New(FizzBuzzRespCacheSize).
			LRU().LoaderFunc(func(key interface{}) (interface{}, error) {
				params := key.(core.FizzBuzzParams)
				return core.FizzBuzz(params.From, params.To, params.Multiple1, params.Multiple2, params.S1, params.S2)
			}).
			Build(),
	}
}

// routes will setup all the available route
func (service *FizzBuzzService) routes() http.Handler {
	service.httpClient = http.NewServeMux()

	service.httpClient.HandleFunc("/fizzbuzz/v1/version", service.VersionRoute)
	service.httpClient.HandleFunc("/fizzbuzz/v1/fizzbuzz", service.FizzBuzzRoute)

	return service.httpClient
}

// Launch FizzBuzz service
func (service *FizzBuzzService) Launch() {
	log.Printf("Listening on %d...\n", service.port)
	http.ListenAndServe(fmt.Sprintf(":%d", service.port), applyCORS(service.routes()))
}

// VersionRoute is a http request handling function which return the version of the service
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

// FizzBuzzRoute is a http request handling function which return the result of the FizzBuzz function
func (service *FizzBuzzService) FizzBuzzRoute(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	to, multiple1, multiple2, word1, word2, err := parseFizzBuzzRouteQueryParam(query)

	if err != nil {
		resp.WriteHeader(http.StatusNotAcceptable)
		data, _ := MessageToJson(CreateErrorMessage("invalid argument", err.Error()))
		resp.Write(data)
		return
	}

	switch req.Method {
	case "GET":
		respond, err := service.fizzBuzzRouteCache.Get(core.FizzBuzzParams{1, to, multiple1, multiple2, word1, word2})
		if err != nil {
			resp.WriteHeader(http.StatusNotAcceptable)
			data, _ := MessageToJson(CreateErrorMessage("invalid argument", err.Error()))
			resp.Write(data)
		}
		data, err := MessageToJson(CreateFizzBuzzMessage(respond.([]string)))
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

func parseFizzBuzzRouteQueryParam(values url.Values) (int, int, int, string, string, error) {

	to, err := strconv.Atoi(values.Get("to"))
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)")
	}
	multiple1, err := strconv.Atoi(values.Get("mul1"))
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)")
	}
	multiple2, err := strconv.Atoi(values.Get("mul2"))
	if err != nil {
		return 0, 0, 0, "", "", fmt.Errorf("One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)")
	}
	word1 := values.Get("word1")
	if len(word1) == 0 {
		return 0, 0, 0, "", "", fmt.Errorf("One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)")
	}
	word2 := values.Get("word2")
	if len(word2) == 0 {
		return 0, 0, 0, "", "", fmt.Errorf("One argument is missing or wrong format. you need to provide: to (int), mul1 (int), mul2 (int), word1 (string), word2 (string)")
	}

	return to, multiple1, multiple2, word1, word2, nil
}
