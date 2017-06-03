package service

import "encoding/json"

const StatusOk = "OK"
const StatusKo = "KO"

type Message interface {
	GetBytesResp() ([]byte, error)
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

func (msg *DefaultMessage) GetBytesResp() ([]byte, error) {
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