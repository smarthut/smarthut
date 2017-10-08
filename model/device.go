package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/smarthut/smarthut/utils"
)

const (
	devicePath = "/data/devices/"
	// dataExt    = ".json"
)

// Device holds abstract device data
type Device struct {
	Host      string    `json:"host"`
	UpdatedAt time.Time `json:"updated_at"`
	Sockets   []Socket  `json:"sockets"`
}

var deviceList map[string]Device

// Socket holds abstract socket data
type Socket struct {
	Value interface{} `json:"value"`
	*SocketInfo
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

func (d *Device) update() error {
	resp, err := http.Get(d.Host)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO: rewrite this part
	var tempDevice deviceAPI

	err = json.Unmarshal(body, &tempDevice)
	if err != nil {
		return err
	}

	d.UpdatedAt = tempDevice.UpdatedAt
	for i := range tempDevice.Sockets {
		d.Sockets[i].Value = tempDevice.Sockets[i]
	}

	return nil
}

// GetDevice returns struct with
func GetDevice(name string) (Device, error) {
	if device, ok := deviceList[name]; ok {
		err := device.update()
		if err != nil {
			return Device{}, err
		}

		return device, nil
	}
	return Device{}, fmt.Errorf("smarthome: no device `%s` found", name)
}

// ListDevices returns device list
func ListDevices() []string {
	names := make([]string, 0, len(deviceList))
	for value := range deviceList {
		names = append(names, value)
	}
	return names
}

// InitializeDevices ...
func InitializeDevices() {
	deviceList = make(map[string]Device)
	deviceNames := utils.ListFilesByExtension(devicePath, dataExt)

	for _, deviceName := range deviceNames {
		d, err := NewDevice(deviceName)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("added device %s with %d sockets", deviceName, len(d.Sockets))
		deviceList[deviceName] = d
	}
}
