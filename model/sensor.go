package model

// Sensor represents basic sensor data
type Sensor struct {
	Value interface{}
	Error error
}

// SensorDescription represents sensor description
type SensorDescription struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Location string `json:"location"`
	*Sensor  `json:"data"`
}
