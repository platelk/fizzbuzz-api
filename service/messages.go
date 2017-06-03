package service

import (
	"encoding/json"
)

// StatusOk is the status code returned on successful request
const StatusOk = 0

// StatusKo is the status code returned on unsuccessful request
const StatusKo = 1

// Message interface define common action which can be made on respond message
type Message interface {
}

// DefaultMessage is a shared struct which contain field required by any return Message
type DefaultMessage struct {
	Status string `json:"status,omitempty"`
	StatusCode uint `json:"statusCode"`
}

// ErrorMessage is return in any error case
type ErrorMessage struct {
	DefaultMessage
	Message string `json:"message,omitempty"`
}

// VersionMessage is return on /version route and contain version of the service
type VersionMessage struct {
	DefaultMessage
	Version string `json:"version"`
}

// FizzBuzzMessage is the respond message on /fizzbuzz route and contain respond
type FizzBuzzMessage struct {
	DefaultMessage
	Response []string
}

// MessageToJson take any message and transform it in a array of by containing JSON format of the Message
func MessageToJson(msg interface{}) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateVersionMessage return a new VersionMessage
func CreateVersionMessage(version string) Message {
	return &VersionMessage{DefaultMessage{"",StatusOk}, version}
}

// CreateErrorMessage return a new ErrorMessage
func CreateErrorMessage(status, message string) Message {
	return &ErrorMessage{DefaultMessage{status, StatusKo}, message}
}

// CreateFizzBuzzMessage return a new FizzBuzzMessage
func CreateFizzBuzzMessage(response []string) Message {
	return &FizzBuzzMessage{DefaultMessage{"", StatusOk}, response}
}