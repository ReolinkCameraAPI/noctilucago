package responses

// A GenericResponse is used to send back information to the client with a general status,
// could be an error or a success message.
// swagger:response genericResponse
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// A VersionResponse is used to send back information to the client with a general status,
// could be an error or a success message.
// swagger:response versionResponse
type VersionResponse struct {
	Version string `json:"version"`
}
