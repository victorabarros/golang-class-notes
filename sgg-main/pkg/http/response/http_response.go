package response

import (
	"net/http"

	"github.com/moshenahmias/failure"
	"github.com/ricardoerikson/sgg/pkg/codes"
)

type HTTPResponse struct {
	statusCode int
	payload    interface{}
}

// Created returns a 201 status code with a payload
func Created(payload interface{}) HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusCreated,
		payload:    payload,
	}
}

// Ok returns a 200 status code with payload
func Ok(payload interface{}) HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusOK,
		payload:    payload,
	}
}

// Deleted returns a standard deleted HTTP response
func Deleted() HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusNoContent,
		payload:    "",
	}
}

// NotFound returns a resource not found response
func NotFound() HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusNotFound,
		payload:    map[string]string{"message": "resource not found"},
	}
}

// BadRequest returns a bad request response
func BadRequest() HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusBadRequest,
		payload:    map[string]string{"message": "check your request before resubmitting"},
	}
}

// UnknownError returns a unknown error
func UnknownError(err error) HTTPResponse {
	return HTTPResponse{
		statusCode: http.StatusInternalServerError,
		payload:    failure.Message(err),
	}
}

// ErrorResponse tries to build an appropriate error message based
// on a list of known errors
func ErrorResponse(err error) HTTPResponse {
	field, e := failure.Field(err, codes.StatusLabel)
	if e != nil {
		return UnknownError(err)
	}

	status, ok := field.(int)
	if !ok {
		return UnknownError(err)
	}

	return HTTPResponse{
		statusCode: status,
		payload:    map[string]string{"message": failure.Message(err)},
	}
}

func CustomResponse(statusCode int, message string) HTTPResponse {
	return HTTPResponse{
		statusCode: statusCode,
		payload:    map[string]string{"message": message},
	}
}
