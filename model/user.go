package model

import (
	"errors"
	"strings"
	"time"

	"github.com/asdine/storm/q"
	"golang.org/x/crypto/bcrypt"

	"github.com/smarthut/smarthut/store"
)

var (
	// ErrNotExist is returned when a user does not exist
	ErrNotExist = errors.New("user does not exist")
	// ErrUserNoName is returned if a new bucket has no name
	ErrUserNoName = errors.New("unable to create an user without a name")
)

// User holds a user data
type User struct {
	ID       int    `json:"id" storm:"id,increment"` // user id
	Username string `json:"username" storm:"unique"` // user name
	Password string `json:"-"`                       // password hash
	Email    string `json:"email" storm:"unique"`
	Name     string `json:"name"`
	Admin    bool   `json:"admin"`
	Role     string `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Credentials holds credential data
type Credentials struct {
	Login    string `json:"login"`    // Username or Email
	Password string `json:"password"` // Password
}

// NewUser initializes a new user from a username, email and password
func NewUser(username, email, password string) (*User, error) {
	if username == "" {
		return nil, ErrUserNoName
	}
	pw, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username:  strings.ToLower(username),
		Email:     email,
		Password:  pw,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, nil
}

// AllUsers returns all registered users
func AllUsers(db *store.DB) ([]User, error) {
	var users []User
	if err := db.All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser finds user by their login (Username or Email)
func GetUser(db *store.DB, login string) (*User, error) {
	var user User
	if err := db.Select(q.Or(
		q.StrictEq("Username", login),
		q.StrictEq("Email", login),
	)).First(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates user information
// TODO: Add proper user fields control
// BODY: Provide proper methods for updating user account fields
func (u *User) Update(db *store.DB, data User) error {
	return db.Update(u)
}

// Delete deletes an user from a database
func (u *User) Delete(db *store.DB) error {
	return db.DeleteStruct(u)
}

// Authenticate a user with a password
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
