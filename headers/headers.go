package headers

// Headers key/value of headers sent with reponse
type Headers map[string]string

// NewHeaders creates a new set of response headers
// -- Contains default headers --
// "access-control-allow-origin":  "*",
// "access-control-allow-headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent",
// "access-control-allow-methods": "OPTIONS,POST,GET,PUT,DELETE",
// "content-type":                 "application/json",
func NewHeaders() *Headers {
	return &Headers{
		"access-control-allow-origin":  "*",
		"access-control-allow-headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent",
		"access-control-allow-methods": "OPTIONS,POST,GET,PUT,DELETE",
		"content-type":                 "application/json",
	}
}

// Add header
func (h *Headers) Add(k, v string) *Headers {
	(*h)[k] = v
	return h
}

// Get header value
func (h *Headers) Get(k string) string {
	return (*h)[k]
}

// GetHeaders returns all headers
func (h *Headers) GetHeaders() Headers {
	headers := *h
	return headers
}
