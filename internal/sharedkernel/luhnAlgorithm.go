package sharedkernel

import (
	"strconv"
	"strings"
)

// Valid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func ValidLuhn(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		num := int(t)

		if double {
			num *= 2
		}

		double = !double

		if num >= 10 {
			num -= 9
		}

		checksum += num
	}

	return checksum
}
