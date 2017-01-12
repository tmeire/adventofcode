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

	password := make([]byte, 0, 8)
	for i := 0; len(password) < 8; i++ {
		h := hash(fmt.Sprintf("ugkcyxxp%d", i))
		if strings.HasPrefix(h, "00000") {
			password = append(password, h[5])
		}
	}
	fmt.Println("password: ", string(password))
}
