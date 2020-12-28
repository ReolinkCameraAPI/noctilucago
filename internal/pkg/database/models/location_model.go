package models

// Location is a physical address.
// Used to map cameras on.
type Location struct {
	// swagger:ignore
	ID uint64 `gorm:"primary_key"`
	// The unique identifier for this Location
	UUID         string `json:"uuid" gorm:"uniqueIndex"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	Province     string `json:"province,omitempty"`
	City         string `json:"city"`
	Country      string `json:"country"`
	// Coordinates
	Latitude  float64 `json:"latitude" gorm:"index"`
	Longitude float32 `json:"longitude" gorm:"index"`
}
