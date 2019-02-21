package model

import "time"

// Channel represent single channel of device
type Channel struct {
	Device      string `json:"device"`
	Socket      int    `json:"socket"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`

	UpdatedAt time.Time
	CreatedAt time.Time
}
