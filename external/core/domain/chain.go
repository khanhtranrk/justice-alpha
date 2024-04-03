package domain

import (
  "time"
  "crypto/md5"
  "encoding/hex"
)

type Chain struct {
  PrimarySafeKey string
  PayloadKey string
  SecondarySafeKey string
  ChainKey string
  RelationshipKey string
  CreatedAt time.Time
}

func NewChain(primarySafeKey string, secondarySafeKey string, payloadKey string, previousChainKey string, relationshipKey string) *Chain {
  hash := md5.Sum([]byte(primarySafeKey + secondarySafeKey + payloadKey + previousChainKey))
  chainKey := hex.EncodeToString(hash[:])
  createdAt := time.Now()

  return &Chain{
    PrimarySafeKey: primarySafeKey,
    SecondarySafeKey: secondarySafeKey,
    PayloadKey: payloadKey,
    ChainKey: chainKey,
    RelationshipKey: relationshipKey,
    CreatedAt: createdAt,
  }
}
