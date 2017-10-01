package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/smarthut/smarthut/utils"
)

const (
	devicePath = "./data/devices/"
	// dataExt    = ".json"
)

// Device holds abstract device data
type Device struct {
	host      string
	UpdatedAt time.Time `json:"updated_at"`
	Sockets   []Socket  `json:"sockets"`
}

var deviceList map[string]Device

// Socket holds abstract socket data
type Socket struct {
	Value []interface{} `json:"value"`
	*SocketInfo
}

// SocketInfo ...
type SocketInfo struct {
	Type     string `json:"type"`     // type string representation, i.e. for providing proper icons
	Location string `json:"location"` // human readable socket description
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

// GetDevice returns struct with
func GetDevice(name string) {

}

// InitializeDevices ...
func InitializeDevices() {
	deviceList = make(map[string]Device)
	deviceNames := utils.ListFilesByExtension(devicePath, dataExt)

	for _, deviceName := range deviceNames {
		d, err := NewDevice(deviceName)
		if err != nil {
			log.Println(err)
		}
		log.Printf("added device %s with %d sockets", deviceName, len(d.Sockets))
		deviceList[deviceName] = d
	}
}
