package util

import (
  "testing"
)

func TestGenerateSafeKey(t *testing.T) {
  key := GenerateSafeKey()

  if key == "" {
    t.Fatalf("GenerateSafeKey(): result is wrong")
  }
}
