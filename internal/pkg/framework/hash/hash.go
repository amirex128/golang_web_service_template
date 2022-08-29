package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/speps/go-hashids"
)

// Sha1 return sha1 hash of a
func Sha1(a string) string {
	h := sha1.New()
	h.Write([]byte(a))
	return fmt.Sprintf("%x", h.Sum(nil))
}

var hids *hashids.HashID

func GetHashids() *hashids.HashID {
	if hids != nil {
		return hids
	}

	hidsData := hashids.NewData()
	hidsData.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	hidsData.Salt = "qtyq68eqeqwy"
	hidsData.MinLength = 6

	h, _ := hashids.NewWithData(hidsData)

	return h
}

func GenerateCode(id int64) string {
	// TODO fix order code hashid generate
	numbers := make([]int, 1)
	numbers[0] = int(id)
	encoded, _ := GetHashids().Encode(numbers)

	return encoded
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
func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
