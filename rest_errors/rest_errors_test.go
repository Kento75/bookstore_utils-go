package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Error())
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("this is the message", errors.New("Causes of error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "rest_error", err.Error)

	assert.NotNil(t, 1, len(err.Causes))
	assert.EqualValues(t, "Causes of error", err.Causes[0])
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("Causes of error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, 1, len(err.Causes))
	assert.EqualValues(t, "Causes of error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}
