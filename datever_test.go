package datever_test

import (
	"testing"

	"github.com/bschaatsbergen/datever"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected *datever.Version
		err      bool
	}{
		{"2024.6.15", &datever.Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, false},
		{"2024.12.31", &datever.Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, false},
		{"2024.2.1-1", &datever.Version{Year: 2024, Month: 2, Day: 1, Patch: "1"}, false},
		{"2024.1.1-alpha", &datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha"}, false},
		{"2024.2.1-alpha001", &datever.Version{Year: 2024, Month: 2, Day: 1, Patch: "alpha001"}, false},
		{"2024.2.1-rc1", &datever.Version{Year: 2024, Month: 2, Day: 1, Patch: "rc1"}, false},
		{"2024.1.1-beta", &datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "beta"}, false},
		{"2024.13.1", nil, true},  // Invalid month
		{"2024.6.15-", nil, true}, // Invalid format
		{"2024.6", nil, true},     // Invalid format
		{"2024.6.1a", nil, true},  // Invalid day format
	}

	for _, test := range tests {
		result, err := datever.ParseVersion(test.input)
		if test.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		version  *datever.Version
		expected string
	}{
		{&datever.Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, "2024.6.15"},
		{&datever.Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, "2024.12.31"},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha"}, "2024.1.1-alpha"},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}, "2024.1.1-alpha001"},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "beta"}, "2024.1.1-beta"},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "rc1"}, "2024.1.1-rc1"},
	}

	for _, test := range tests {
		result := test.version.String()
		assert.Equal(t, test.expected, result)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		version  *datever.Version
		expected bool
	}{
		{&datever.Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, true},
		{&datever.Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, true},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "1"}, true},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "rc1"}, true},
		{&datever.Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}, true},
		{&datever.Version{Year: 2024, Month: 13, Day: 15, Patch: ""}, false}, // Invalid month
		{&datever.Version{Year: 2024, Month: 0, Day: 15, Patch: ""}, false},  // Invalid month
		{&datever.Version{Year: 2024, Month: 6, Day: 0, Patch: ""}, false},   // Invalid day
	}

	for _, test := range tests {
		result := test.version.IsValid()
		assert.Equal(t, test.expected, result)
	}
}
