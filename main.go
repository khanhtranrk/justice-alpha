package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	// publicKey := privateKey.PublicKey

  msg := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
  hash := sha256.New().Sum(msg)
  reader := bytes.NewReader(msg)

  ecdsa.Sign(reader, privateKey, hash)
  privateKey.
}
