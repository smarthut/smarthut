package model

import (
	"time"

	"github.com/smarthut/smarthut/store"
)

const (
	devicePath = "/data/devices/"
	// dataExt    = ".json"
)

// Device holds device data
type Device struct {
	ID      int          `json:"id" storm:"id,increment"` // device id
	Name    string       `json:"name" storm:"unique"`     // device slug
	Host    string       `json:"host" storm:"unique"`     // device url
	Title   string       `json:"title"`
	Sockets []SocketInfo `json:"sockets"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var deviceList map[string]Device

// Socket holds abstract socket data
type Socket struct {
	Value interface{} `json:"value"`
	// *SocketInfo
}

// SocketInfo ...
type SocketInfo struct {
	Type     string `json:"type"`     // type string representation, i.e. for providing proper icons
	Location string `json:"location"` // human readable socket description
}

type deviceAPI struct {
	UpdatedAt time.Time     `json:"updated_at"`
	Sockets   []interface{} `json:"sockets"`
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

// SetSockets ...
func (d *Device) SetSockets(db *store.DB, sockets []SocketInfo) error {
	d.Sockets = sockets
	d.UpdatedAt = time.Now()
	return db.UpdateField(d, "Sockets", d.Sockets)
}
