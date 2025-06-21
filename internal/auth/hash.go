package auth

import (
	"crypto/sha256"
	"fmt"
)

// HashToken crea un hash SHA-256 de un string de token.
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", hash)
}
