package rest_errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Error())
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
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
