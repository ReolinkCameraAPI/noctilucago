package models

// ProxyInput is a network configuration for connecting to a proxy server
// The ProxyInput setting can be reused across many cameras
// swagger:model
type ProxyInput struct {
	// Host is the IP or Domain to connect to
	// required: true
	// example: 192.168.1.1 or example.com
	Host string `json:"host" binding:"required"`

	// Port is used together with Host. If none are specified, no Port will be used.
	// required: false
	// example: 9000
	Port string `json:"port,omitempty"`

	// Username for authenticating with the proxy server.
	// required: false
	Username string `json:"username,omitempty"`

	// Password for authenticating with the proxy server.
	// required: false
	Password string `json:"password,omitempty"`

	// Scheme for connecting to the ProxyInput server, such as HTTP, HTTPS or SOCKS5
	// required: true
	// example: HTTPS
	Scheme string `json:"scheme" binding:"required"`

	// Protocol to use, for connecting with the ProxyInput server, such as TCP or UDP
	// required: true
	// example: UDP
	Protocol string `json:"protocol" binding:"required"`
}