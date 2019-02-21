package model

import (
	"errors"
	"strings"
	"time"

	"github.com/smarthut/smarthut/store"
)

var (
	// ErrDeviceNoName is returned if a new device has no name
	ErrDeviceNoName = errors.New("unable to create a bucket without a name")
)

// Device holds device data
type Device struct {
	ID          int    `json:"id" storm:"id,increment"` // device id
	Name        string `json:"name" storm:"unique"`     // device slug
	Host        string `json:"host" storm:"unique"`     // device url
	Title       string `json:"title"`
	Description string `json:"description"`
	Sockets     []struct {
		Type        string `json:"type"`        // type string representation, i.e. for providing proper icons
		Description string `json:"description"` // human readable socket description
	} `json:"sockets"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewDevice creates a new device
func NewDevice(name, host, title string) (*Device, error) {
	if name == "" {
		return nil, ErrBucketNoName
	}
	return &Device{
		Name:      strings.ToLower(name),
		Host:      host,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// AllDevices retursa all available devices list
func AllDevices(db *store.DB) ([]Device, error) {
	var devices []Device
	if err := db.All(&devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// GetDevice finds a device by name
func GetDevice(db *store.DB, name string) (*Device, error) {
	var device Device
	if err := db.One("Name", name, &device); err != nil {
		return nil, err
	}
	return &device, nil
}

// Delete deletes a device from a database
func (d *Device) Delete(db *store.DB) error {
	return db.DeleteStruct(d)
}
