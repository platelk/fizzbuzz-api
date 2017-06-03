package http

// HttpService define basic possible interaction for all HttpService
type HttpService interface {
	Launch()
}

// FizzBuzzService is a HttpService which
type FizzBuzzService struct {

}

func CreateFizzBuzzService() HttpService {
	return &FizzBuzzService{}
}

func (service *FizzBuzzService) Launch() {

}