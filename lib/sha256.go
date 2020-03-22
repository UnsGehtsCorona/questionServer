package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func Sha256(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateUid(s string) string {
	random := fmt.Sprintf("%d", rand.Intn(100000))
	return Sha256(time.Now().String() + s + random)
}
