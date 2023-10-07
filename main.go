package main

import (
	"github.com/andrejrakic/dapp-development-with-go/blocks"
	"github.com/andrejrakic/dapp-development-with-go/client"
)

func main() {
	c := client.Connect()
	Address()

	// get latest block details
	blocks.LatestBlockInfo(c)
}
