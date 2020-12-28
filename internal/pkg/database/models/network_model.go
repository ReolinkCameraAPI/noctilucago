package models

// Proxy is a network configuration for connecting to a proxy server
type Proxy struct {
	// swagger:ignore
	ID       uint64 `gorm:"primary_key"`
	UUID     string `json:"uuid,omitempty"`
	Host     string `json:"host"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Protocol string `json:"protocol"`
	Scheme   string `json:"scheme"`
}
