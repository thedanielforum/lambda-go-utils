package events

// APIGatewayProxyRequest contains data coming from the API Gateway proxy
type APIGatewayProxyRequest struct {
	Resource              string                        `json:"resource"` // The resource path defined in API Gateway
	Path                  string                        `json:"path"`     // The url path for the caller
	HTTPMethod            string                        `json:"httpMethod"`
	Headers               map[string]string             `json:"headers"`
	QueryStringParameters map[string]string             `json:"queryStringParameters"`
	PathParameters        map[string]string             `json:"pathParameters"`
	StageVariables        map[string]string             `json:"stageVariables"`
	RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
	Body                  string                        `json:"body"`
	IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyRequestContext contains the information to identify the AWS account and resources invoking the
// Lambda function. It also includes Cognito identity information for the caller.
type APIGatewayProxyRequestContext struct {
	AccountID    string                    `json:"accountId"`
	ResourceID   string                    `json:"resourceId"`
	Stage        string                    `json:"stage"`
	RequestID    string                    `json:"requestId"`
	Identity     APIGatewayRequestIdentity `json:"identity"`
	ResourcePath string                    `json:"resourcePath"`
	HTTPMethod   string                    `json:"httpMethod"`
	APIID        string                    `json:"apiId"` // The API Gateway rest API Id
	Authorizer   AuthClaims                `json:"authorizer"`
}

// AuthClaims jwt claims
type AuthClaims struct {
	JWTClaims Claims `json:"claims"`
}

// Claims jwt claims
type Claims struct {
	Sub             string `json:"sub"`
	Aud             string `json:"aud"`
	Exp             string `json:"exp"`
	Iss             string `json:"iss"`
	TokenUse        string `json:"token_use"`
	Email           string `json:"email"`
	EmailVerified   string `json:"email_verified"`
	CognitoUsername string `json:"cognito:username"`
	EventID         string `json:"event_id"`
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId"`
	AccountID                     string `json:"accountId"`
	CognitoIdentityID             string `json:"cognitoIdentityId"`
	Caller                        string `json:"caller"`
	APIKey                        string `json:"apiKey"`
	SourceIP                      string `json:"sourceIp"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider"`
	UserArn                       string `json:"userArn"`
	UserAgent                     string `json:"userAgent"`
	User                          string `json:"user"`
}

// APIGatewayCustomAuthorizerContext represents the expected format of an API Gateway custom authorizer response.
type APIGatewayCustomAuthorizerContext struct {
	PrincipalID *string `json:"principalId"`
	StringKey   *string `json:"stringKey,omitempty"`
	NumKey      *int    `json:"numKey,omitempty"`
	BoolKey     *bool   `json:"boolKey,omitempty"`
}
