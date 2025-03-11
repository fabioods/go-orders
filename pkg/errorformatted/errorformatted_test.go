package errorformated

import (
	"errors"
	"net/http"
	"testing"

	"github.com/fabioods/go-orders/pkg/trace"
	"github.com/stretchr/testify/assert"
)

func TestErrorFormatted(t *testing.T) {
	errorF := ErrorFormated{
		Code:    "error-code",
		Message: "error-message",
		Status:  http.StatusBadRequest,
		Cause:   errors.New("error-cause"),
		Trace:   trace.GetTrace(),
	}
	assert.Equal(t, errorF.Error(), "error-message")
	assert.Equal(t, http.StatusBadRequest, errorF.StatusCode())
}

func TestErrorFormatted_NoStatusCode(t *testing.T) {
	errorF := ErrorFormated{
		Code:    "error-code",
		Message: "error-message",
		Cause:   errors.New("error-cause"),
		Trace:   trace.GetTrace(),
	}
	assert.Equal(t, errorF.Error(), "error-message")
	assert.Equal(t, http.StatusInternalServerError, errorF.StatusCode())
}

func TestErrorFormatted_BadRequestError(t *testing.T) {
	errorFormatted := BadRequestError(trace.GetTrace(), "error-message", "error-cause")
	assert.Equal(t, errorFormatted.Error(), "error-cause")
	assert.Equal(t, http.StatusBadRequest, errorFormatted.StatusCode())
}

func TestErrorFormatted_NotFoundError(t *testing.T) {
	errorFormated := NotFoundError(trace.GetTrace(), "error-message", "error-cause")
	assert.Equal(t, errorFormated.Error(), "error-cause")
	assert.Equal(t, http.StatusNotFound, errorFormated.StatusCode())
}

func TestErrorFormatted_UnexpectedError(t *testing.T) {
	errorFormated := UnexpectedError(trace.GetTrace(), "error-message", "error-cause")
	assert.Equal(t, errorFormated.Error(), "error-cause")
	assert.Equal(t, http.StatusInternalServerError, errorFormated.StatusCode())
}

func TestErrorFormatted_UnprocesableEntityError(t *testing.T) {
	errorFormated := UnprocesableEntityError(trace.GetTrace(), "error-message", "error-cause")
	assert.Equal(t, errorFormated.Error(), "error-cause")
	assert.Equal(t, http.StatusUnprocessableEntity, errorFormated.StatusCode())
}
