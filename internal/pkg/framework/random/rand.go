package random

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

// ID random generator
var ID = make(chan string)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
	go func() {
		h := sha1.New()
		c := []byte(time.Now().String())
		for {
			_, _ = h.Write(c)
			ID <- fmt.Sprintf("%x", h.Sum(nil))
		}
	}()
}
