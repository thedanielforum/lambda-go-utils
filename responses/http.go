package responses

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/events"
)

// HTTPResponse formats http response
func HTTPResponse(status int, body interface{}, headers Headers) events.APIGatewayProxyResponse {
	// JSON encode body
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.WithError(err).Error("failed to Marshal response body")
		// return error instead
		return HTTPError(http.StatusInternalServerError, err, headers)
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

func HTTPError(status int, err error, headers Headers) events.APIGatewayProxyResponse {
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
