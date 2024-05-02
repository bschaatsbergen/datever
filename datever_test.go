package datever

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected *Version
		err      bool
	}{
		{"v2024.6.15", &Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, false},
		{"v2024.12.31", &Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, false},
		{"v2024.2.1-1", &Version{Year: 2024, Month: 2, Day: 1, Patch: "1"}, false},
		{"v2024.1.1-alpha", &Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha"}, false},
		{"v2024.2.1-alpha001", &Version{Year: 2024, Month: 2, Day: 1, Patch: "alpha001"}, false},
		{"v2024.2.1-rc1", &Version{Year: 2024, Month: 2, Day: 1, Patch: "rc1"}, false},
		{"v2024.1.1-beta", &Version{Year: 2024, Month: 1, Day: 1, Patch: "beta"}, false},
		{"v2024.13.1", nil, true},  // Invalid month
		{"v2024.6.15-", nil, true}, // Invalid format
		{"v2024.6", nil, true},     // Invalid format
		{"v2024.6.1a", nil, true},  // Invalid day format
	}

	for _, test := range tests {
		result, err := ParseVersion(test.input)
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
		version  *Version
		expected string
	}{
		{&Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, "v2024.6.15"},
		{&Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, "v2024.12.31"},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha"}, "v2024.1.1-alpha"},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}, "v2024.1.1-alpha001"},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "beta"}, "v2024.1.1-beta"},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "rc1"}, "v2024.1.1-rc1"},
	}

	for _, test := range tests {
		result := test.version.String()
		assert.Equal(t, test.expected, result)
	}
}

func TestCompare(t *testing.T) {
	v1 := &Version{Year: 2024, Month: 6, Day: 15, Patch: ""}
	v2 := &Version{Year: 2024, Month: 12, Day: 31, Patch: ""}
	v3 := &Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}
	v4 := &Version{Year: 2024, Month: 1, Day: 1, Patch: "beta"}
	v5 := &Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}

	tests := []struct {
		v1, v2   *Version
		expected int
	}{
		{v1, v2, -1},
		{v2, v1, 1},
		{v3, v4, -1},
		{v4, v3, 1},
		{v3, v5, 0},
	}

	for _, test := range tests {
		result := test.v1.Compare(test.v2)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		version  *Version
		expected bool
	}{
		{&Version{Year: 2024, Month: 6, Day: 15, Patch: ""}, true},
		{&Version{Year: 2024, Month: 12, Day: 31, Patch: ""}, true},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "1"}, true},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "rc1"}, true},
		{&Version{Year: 2024, Month: 1, Day: 1, Patch: "alpha001"}, true},
		{&Version{Year: 2024, Month: 13, Day: 15, Patch: ""}, false}, // Invalid month
		{&Version{Year: 2024, Month: 0, Day: 15, Patch: ""}, false},  // Invalid month
		{&Version{Year: 2024, Month: 6, Day: 0, Patch: ""}, false},   // Invalid day
	}

	for _, test := range tests {
		result := test.version.IsValid()
		assert.Equal(t, test.expected, result)
	}
}
