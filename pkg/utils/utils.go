package utils

import (
	"crypto/sha256"
	"fmt"
)

// Hash Adapted from: https://blog.8bitzen.com/posts/22-08-2019-how-to-hash-a-struct-in-go
func Hash(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
