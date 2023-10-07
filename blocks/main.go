package blocks

import (
	"context"
	// "encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func LatestBlockInfo(client *ethclient.Client) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	latestBlockNumber := header.Number

	fmt.Println("The latest block number is:", latestBlockNumber.String())

	fmt.Println("The latest block details are: ")

	latestBlock, _ := client.BlockByNumber(context.Background(), latestBlockNumber)

	fmt.Println("Block number: ", latestBlock.Number().Uint64())
	fmt.Println("Block timestamp: ", latestBlock.Time())
	fmt.Println("Block dificulty: ", latestBlock.Difficulty().Uint64())
	fmt.Println("Block hash: ", latestBlock.Hash().Hex())
	fmt.Println("Number of transactions", len(latestBlock.Transactions()))

	fmt.Println("Transaction details: ")

	for i := 0; i < len(latestBlock.Transactions()); i++ {
		currentTx := latestBlock.Transactions()[i]
		fmt.Println("--------------------------------------------")
		fmt.Println("Current transaction hash: ", currentTx.Hash())
		// fmt.Println("Current transaction data: ", hex.EncodeToString(currentTx.Data()))
		fmt.Println("Current transaction access list: 0x", currentTx.AccessList())
		fmt.Println("Current transaction gas: ", currentTx.Gas())
		fmt.Println("Current transaction size: ", currentTx.Size())
		fmt.Println("Current transaction value: ", currentTx.Value())
	}
}
