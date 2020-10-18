package utils

import (
	"fmt"
	"regexp"
)

const (
	intPattern  string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	floatPatern string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"

	invalidValue  string = "%s value: %s not valid number"
	invalidValues string = "%s value: %s not valid number, %s value: %s not valid number"
)

var (
	rxInt   = regexp.MustCompile(intPattern)
	rxFloat = regexp.MustCompile(floatPatern)
)

// IsXYValid checks weather x and y are valid integer or float values.
func IsXYValid(x, y string) (bool, error) {
	validX := isIntOrFloat(x)
	validY := isIntOrFloat(y)

	if !validX && !validY {
		return false, fmt.Errorf(invalidValues, "x", x, "y", y)
	} else if !validX {
		return false, fmt.Errorf(invalidValue, "x", x)
	} else if !validY {
		return false, fmt.Errorf(invalidValue, "y", y)
	}

	return true, nil
}

func isIntOrFloat(str string) bool {
	return str != "" && (rxInt.MatchString(str) || rxFloat.MatchString(str))
}
