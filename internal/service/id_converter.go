package service

import (
	"fmt"
	"strconv"
)

func ConvertIdToHex(id uint64) string {
	return fmt.Sprintf("%x", id)
}

func ConvertHexToId(hexId string) (int, error) {
	number, err := strconv.ParseUint(hexId, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hex string: %w", err)
	}
	return int(number), nil
}
