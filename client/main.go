package client

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func Connect() *ethclient.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyKey := os.Getenv("ALCHEMY_KEY")
	sepoliaRpcUrl := fmt.Sprintf("https://eth-sepolia.g.alchemy.com/v2/%s", alchemyKey)

	client, err := ethclient.Dial(sepoliaRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	return client // we'll use this in the upcoming sections
}
