package model

import (
	"time"
)

// Device holds device data
type Device struct {
	ID      int    `json:"id" storm:"id,increment"` // device id
	Name    string `json:"name" storm:"unique"`     // device slug
	Host    string `json:"host" storm:"unique"`     // device url
	Title   string `json:"title"`
	Sockets []struct {
		Type        string `json:"type"`        // type string representation, i.e. for providing proper icons
		Description string `json:"description"` // human readable socket description
	} `json:"sockets"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewDevice ...
func NewDevice(name, host, title string) (*Device, error) {
	device := &Device{
		Name:      name,
		Host:      host,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return device, nil
}
