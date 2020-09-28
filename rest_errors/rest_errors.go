package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusNotImplemented,
		Error:   "rest_error",
	}

	// エラー原因があるなら追加
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	// エラー原因があるなら追加
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
