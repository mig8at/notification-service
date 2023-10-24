package pkg

import (
	"fmt"
	"strconv"
)

func ConvertToSeconds(input string) (int64, error) {
	if len(input) < 2 {
		return 0, fmt.Errorf("input string too short")
	}

	unit := input[len(input)-1:]
	valueStr := input[:len(input)-1]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}

	switch unit {
	case "s":
		return int64(value), nil
	case "m":
		return int64(value) * 60, nil
	case "h":
		return int64(value) * 60 * 60, nil
	case "d":
		return int64(value) * 60 * 60 * 24, nil
	default:
		return 0, fmt.Errorf("unknown time unit: %s", unit)
	}
}
