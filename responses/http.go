package responses

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/thedanielforum/lambda-go-utils/events"
	"github.com/thedanielforum/lambda-go-utils/headers"
)

// OK formats http response
func OK(status int, body interface{}, headers headers.Headers) events.APIGatewayProxyResponse {
	// JSON encode body
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.WithError(err).Error("failed to Marshal response body")
		// return error instead
		return Error(http.StatusInternalServerError, err, headers)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBody),
		Headers:    headers,
		StatusCode: status,
	}
}

// ErrMsg error message struct for json response formating
type ErrMsg struct {
	Message string `json:"message"`
}

// Error response
func Error(status int, err error, headers headers.Headers) events.APIGatewayProxyResponse {
	msg := "Internal server error"
	if err != nil {
		msg = err.Error()
	}

	jsonBody, err := json.Marshal(&ErrMsg{
		Message: msg,
	})
	if err != nil {
		log.WithError(err).Error("failed to Marshal error response body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBody),
		Headers:    headers,
		StatusCode: status,
	}
}
