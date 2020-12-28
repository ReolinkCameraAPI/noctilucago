package responses

// A SessionResponse is used to send back jwt tokens to the client.
// swagger:response sessionResponse
type SessionResponse struct {
	Token string `json:"token"`
}
