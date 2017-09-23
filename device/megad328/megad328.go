package megad328

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURL    string        = "http://%s/%s/?pt=%s&cmd=get"
	sensorsNum int           = 14
	timeout    time.Duration = 5 * time.Second
)

// MegaD328 holds MegaD328 information
type MegaD328 struct {
	url     string
	client  http.Client
	Sensors `json:"sensors"`
}

// Sensor represents a sensor value
type Sensor struct {
	Value interface{}
	Error error
}

// Sensors holds all sensor value
type Sensors [sensorsNum]Sensor

// New creates new MegaD328 connection
func New(host, password string) MegaD328 {
	url := fmt.Sprintf(baseURL, host, password, "%d")
	d := MegaD328{
		url:    url,
		client: http.Client{Timeout: timeout},
	}
	return d
}

// Sensor returns value of sensor by it's ID
func (d MegaD328) Sensor(id int) (value interface{}, err error) {
	url := fmt.Sprintf(d.url, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value, err = strconv.ParseFloat(string(body), 64)
	if err != nil {
		// unable to parse fload, return human-readable value
		return string(body), nil
	}
	return
}
