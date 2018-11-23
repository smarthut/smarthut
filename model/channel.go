package model

import "time"

// Channel represent single channel of device
type Channel struct {
	Device      string `toml:"device"`
	Socket      int
	Type        string
	Title       string
	Description string
	Icon        string

	UpdatedAt time.Time
	CreatedAt time.Time
}
