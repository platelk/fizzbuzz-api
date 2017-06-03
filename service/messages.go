package service

import (
	"encoding/json"
)

const StatusOk = "OK"
const StatusKo = "KO"

type Message interface {
}

type DefaultMessage struct {
	Status string `json:"status,omitempty"`
	StatusCode string `json:"statusCode"`
}

type ErrorMessage struct {
	DefaultMessage
	Message string `json:"message,omitempty"`
}

type VersionMessage struct {
	DefaultMessage
	Version string `json:"version"`
}

type FizzBuzzMessage struct {
	DefaultMessage
	Response []string
}

func MessageToJson(msg interface{}) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateVersionMessage(version string) Message {
	return &VersionMessage{DefaultMessage{"",StatusOk}, version}
}

func CreateErrorMessage(status, message string) Message {
	return &ErrorMessage{DefaultMessage{status, StatusKo}, message}
}

func CreateFizzBuzzMessage(response []string) Message {
	return &FizzBuzzMessage{DefaultMessage{"", StatusOk}, response}
}