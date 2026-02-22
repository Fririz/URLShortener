package service

import (
	"fmt"
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ConvertIdToBase62(id uint64) string {
	if id == 0 {
		return string(base62Chars[0])
	}

	var b [11]byte
	i := len(b)

	for id > 0 {
		i--
		b[i] = base62Chars[id%62]
		id /= 62
	}

	return string(b[i:])
}


func ConvertBase62ToId(base62 string) (int, error) {
	var id uint64

	for _, char := range base62 {
		index := strings.IndexRune(base62Chars, char)
		if index == -1 {
			return 0, fmt.Errorf("invalid base62 string: invalid character '%c'", char)
		}
		
		id = id*62 + uint64(index)
	}

	return int(id), nil
}