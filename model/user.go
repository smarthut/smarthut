package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/smarthut/smarthut/utils"
)

const (
	userPath = "./data/users/"
	dataExt  = ".json"
)

var (
	// ErrNotExist is returned when a user does not exist
	ErrNotExist = errors.New("user does not exist")
)

// User holds a user data
type User struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Password string `json:"password"` // omitted from API
	Email    string `json:"email"`
}

// ListUsers all users
func ListUsers() []string {
	return utils.ListFilesByExtension(userPath, dataExt)
}

// GetUser by it's login (filename)
func GetUser(login string) (User, error) {
	path := userPath + login + dataExt

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
