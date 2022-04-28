package instrumented

import (
	"crypto/rand"
	"encoding/hex"
)

func generateID() string {
	random := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := rand.Read(random); err != nil {
		panic(err)
	}
	return hex.EncodeToString(random)
}
