package util

import (
  "crypto/md5"
  "encoding/hex"
  "github.com/google/uuid"
)

func GenerateRelationshipKey() string {
  hash := md5.Sum([]byte(uuid.New().NodeID()))
  return hex.EncodeToString(hash[:])
}
