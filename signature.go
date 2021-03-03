package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func getSignature(query string) string {
	key := []byte(os.Getenv("SECRET_KEY"))
	h := hmac.New(sha256.New, key)
	h.Write([]byte(query))
	return hex.EncodeToString(h.Sum(nil))
}
