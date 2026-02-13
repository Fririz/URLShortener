package service

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func ConvertIdToHex(id int) string {
	number := uint64(id)
	bytesContainer := make([]byte, 8)
	binary.BigEndian.PutUint64(bytesContainer, number)
	return hex.EncodeToString(bytesContainer)
}

func ConvertHexToId(hexId string) (int, error) {
	bytesContainer, err := hex.DecodeString(hexId)
	if err != nil {
		return 0, fmt.Errorf("invalid hex string: %w", err)
	}

	if len(bytesContainer) != 8 {
		return 0, fmt.Errorf("invalid length: expected 8 bytes")
	}
	number := binary.BigEndian.Uint64(bytesContainer)

	return int(number), nil
}
