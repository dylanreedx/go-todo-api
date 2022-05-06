package tests

import (
	"testing"

	"github.com/carbondesigned/go-todo/utils"
	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		description string
		password    string
		hash        string
		expected    bool
	}{
		{
			description: "successfully dehash a password",
			password:    "password",
			hash:        "$2a$10$e7grJwxMeTnjcB4SswjaiO5N6MO2tt5J02b173iI2rtn7W6rLciGe",
			expected:    true,
		},
	}
	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		assert.Equalf(t, test.expected, utils.CheckPasswordHash(test.password, test.hash), test.description)
	}
}
