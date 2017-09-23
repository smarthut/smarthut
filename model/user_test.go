package model

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type testpair struct {
	value  string
	result bool
}

var tests = []testpair{
	{"myP@44w0rd!", true},
	{"Not-myP@44w0rd!", false},
}

func TestValidate(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("myP@44w0rd!"), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}

	u := User{Password: string(hashedPassword)}

	for _, pair := range tests {
		err := u.Validate(pair.value)
		v := err == nil
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
