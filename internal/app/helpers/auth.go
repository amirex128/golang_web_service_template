package helpers

import (
	"backend/internal/pkg/framework/hash"
)

func GeneratePasswordHash(pass string) string {
	return hash.Sha512EncodeSaltIter(pass, 2, "amirex128-hamyarsale")
}
