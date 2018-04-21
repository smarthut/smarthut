package model

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/smarthut/smarthut/store"
)

const (
	userPath = "/data/users/"
	dataExt  = ".json"
)

var (
	// ErrNotExist is returned when a user does not exist
	ErrNotExist = errors.New("user does not exist")
)

// User holds a user data
type User struct {
	Login    string `json:"login" storm:"id"` // user name
	Password string `json:"-"`                // encrypted password
	Email    string `json:"email" storm:"unique"`
	Name     string `json:"name"`
	Role     string `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser initializes a new user from a login, email and password
func NewUser(login, email, password string) (*User, error) {
	pw, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Login:     login,
		Email:     email,
		Password:  pw,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, nil
}

// SetRole sets the users Role with roleName
func (u *User) SetRole(db *store.DB, roleName string) error {
	u.Role = strings.TrimSpace(roleName)
	u.UpdatedAt = time.Now()
	return db.UpdateField(u, "Role", u.Role)
}

// Authenticate a user from a password
func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}
