package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	password := make([]byte, 8)

	matches := 0
	for i := 0; matches < 8; i++ {
		h := hash(fmt.Sprintf("ugkcyxxp%d", i))
		if strings.HasPrefix(h, "00000") {
			idx := int(h[5]) - 48
			if idx < 0 || idx >= 8 {
				continue
			}

			if password[idx] == byte(0) {
				password[idx] = h[6]
				matches++
			}
		}
	}
	fmt.Println("password: ", string(password))
}
