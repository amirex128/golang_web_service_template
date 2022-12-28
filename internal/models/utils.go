package models

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func GeneratePasswordHash(pass string) string {
	return Sha512EncodeSaltIter(pass, 2, "amirex128-selloora")
}
func Sha512EncodeSaltIter(raw string, iter int, salt string) string {
	salted := fmt.Sprintf("%s{%s}", raw, salt)
	digest := mSha512(salted)
	for i := 1; i < iter; i++ {
		digest = mSha512(string(digest) + salted)
	}
	return hex.EncodeToString(digest)
}

func mSha512(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
