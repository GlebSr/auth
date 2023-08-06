package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_EncryptPassword(t *testing.T) {
	tests := []struct {
		name    string
		user    func() *User
		isValid bool
	}{
		{
			name: "Valid",
			user: func() *User {
				return NewUser(t)
			},
			isValid: true,
		},
		{
			name: "empty password",
			user: func() *User {
				us := NewUser(t)
				us.Password = ""
				return us
			},
			isValid: false,
		},
		{
			name: "long password",
			user: func() *User {
				us := NewUser(t)
				us.Password = "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
				return us
			},
			isValid: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.NoError(t, test.user().EncryptPassword())
			} else {
				assert.Error(t, test.user().EncryptPassword())
			}
		})
	}

}
