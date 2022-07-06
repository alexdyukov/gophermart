package sharedkernel

import (
	"strconv"
	"strings"
)

func ValidLuhn(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, err := strconv.ParseInt(source[i], 10, 8)
		if err != nil {
			return 1
		}
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
