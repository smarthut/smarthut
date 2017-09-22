package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	dataPath = "./data/users/"
	dataExt  = ".json"
)

var (
	// ErrNotExist is returned when a user does not exist
	ErrNotExist = errors.New("user does not exist")
)

// User holds a user data
// TODO: make password private variable
type User struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ListUsers all users
func ListUsers() []string {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	var result []string

	for _, f := range files {
		// Do not expose any non-JSON files
		basename := f.Name()
		if path.Ext(basename) == dataExt {
			basename = strings.TrimSuffix(basename, filepath.Ext(basename))
			result = append(result, basename)
		}
	}

	return result
}

// GetUser by it's login (filename)
func GetUser(login string) (User, error) {

	path := dataPath + login + dataExt

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return User{}, ErrNotExist
	}

	// Read related json file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return User{}, err
	}

	var u User
	err = json.Unmarshal(file, &u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

// Validate checks if provided password is valid
func (u User) Validate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
