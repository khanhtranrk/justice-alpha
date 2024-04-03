package main

import (
	"fmt"

	"github.com/khanhtranrk/justice-alpha/external/chain"
)

func main() {
  primarySafeKey := "0a9c2f825cd7ce73bf3d57857813ba5f"
  secondarySafeKey := "bbe98da28d252fe20b8b7ebdebb1f056"
  payloadKey := "c07fba0b475ae107976757de05446f61"
  previousChainKey := "728fd79b17a353bd90dfc255711864e6"

  fmt.Println(chain.NewChain(primarySafeKey, secondarySafeKey, payloadKey, previousChainKey))
}
