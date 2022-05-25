package cryptx

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, pass string) string {
	//AES-256 - 32 bits key ,from document example
	bs, _ := scrypt.Key([]byte(pass), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%v", bs)
}
