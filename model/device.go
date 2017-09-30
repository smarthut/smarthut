package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	devicePath = "./data/devices/"
	// dataExt    = ".json"
)

// Device holds abstract device data
type Device struct {
	Name     string   `json:"name"`
	Driver   string   `json:"driver"`
	Host     string   `json:"host"`
	Password string   `json:"password,omitempty"`
	Sockets  []Socket `json:"sockets"`
}

// Driver defines abstract device driver
type Driver interface {
	New()
	Name()
	Get()
	Set()
}

// Socket holds abstract socket data
type Socket struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`        // type string representation, i.e. for providing proper icons
	Description string `json:"description"` // human readable socket description
}

// NewDevice creates new device
func NewDevice(id string) (Device, error) {
	path := devicePath + id + dataExt

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Device{}, ErrNotExist
	}

	// Read related json file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Device{}, err
	}

	var d Device
	err = json.Unmarshal(file, &d)
	if err != nil {
		return Device{}, err
	}

	return d, nil
}

func init() {
}
