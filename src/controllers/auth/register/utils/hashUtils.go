package registerUtils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashCID(cid string) string {
	hash := sha256.Sum256([]byte(cid))
	return hex.EncodeToString(hash[:])
}
