package helpers

import (
	"errors"

	"github.com/thedanielforum/lambda-go-utils/events"
)

// AuthorizeUser exstracts and checks users claims from JWT token
func AuthorizeUser(request events.APIGatewayProxyRequest) (events.Claims, error) {
	userID := request.RequestContext.Authorizer.JWTClaims.CognitoUsername
	if userID == "" {
		return events.Claims{}, errors.New("Failed to authenticate user")
	}

	return request.RequestContext.Authorizer.JWTClaims, nil
}
