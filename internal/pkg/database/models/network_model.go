package models

// Proxy is a network configuration for connecting to a proxy server
// swagger:model
type Proxy struct {
	ID       uint64 `gorm:"primary_key"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
}
