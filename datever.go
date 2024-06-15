package datever

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Version represents a date-based version with an optional patch.
type Version struct {
	Year, Month, Day int
	Patch            string
}

// ParseVersion parses a version string into a Version struct.
// Supported formats are:
// * YYYY.M.D
// * YYYY.M.D-PATCH
// where PATCH can be any alphanumeric string.
func ParseVersion(version string) (*Version, error) {
	// Regex pattern to match versions
	pattern := `^(\d{4})\.(\d{1,2})\.(\d{1,2})(?:-([a-zA-Z0-9]+))?$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(version)

	if matches == nil {
		return nil, errors.New("invalid version format")
	}

	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid year: %w", err)
	}
	month, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid month: %w", err)
	}
	day, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid day: %w", err)
	}

	patch := matches[4]

	dateVersion := &Version{
		Year:  year,
		Month: month,
		Day:   day,
		Patch: patch,
	}

	if !dateVersion.IsValid() {
		return nil, errors.New("invalid version: date components are out of range")
	}

	return dateVersion, nil
}

// String returns the string representation of the Version.
func (v *Version) String() string {
	if v.Patch != "" {
		return fmt.Sprintf("%d.%d.%d-%s", v.Year, v.Month, v.Day, v.Patch)
	}
	return fmt.Sprintf("%d.%d.%d", v.Year, v.Month, v.Day)
}

// IsValid checks if the Version is valid.
func (v *Version) IsValid() bool {
	return v.Year > 0 && v.Month > 0 && v.Month <= 12 && v.Day > 0 && v.Day <= 31
}
