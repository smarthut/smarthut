package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// JWTConfiguration holds all the JWT related configuration
type JWTConfiguration struct {
	Secret string `json:"secret" requred:"true"`
}

// Configuration holds configuration struct
type Configuration struct {
	API struct {
		Host string
		Port int `envconfig:"PORT" default:"8080"`
	}
	JWT JWTConfiguration `json:"jwt"`
}

func loadEnvironment(filename string) error {
	if filename != "" {
		return godotenv.Load(filename)
	}
	err := godotenv.Load()
	// handle error is .env does not exist
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// Load loads configuration
func Load(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("smarthut", config); err != nil {
		return nil, err
	}

	return config, nil
}
