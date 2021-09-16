package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wo4zhuzi/solana-go-sdk/client"
	"github.com/wo4zhuzi/solana-go-sdk/client/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, decimals, err := c.GetTokenAccountBalance(
		context.Background(),
		"AyHWro8zumyZN68Mhuk6mhNUUQ2VX5qux2pMD4HnN3aJ",
	)
	if err != nil {
		log.Fatalln("get balance error", err)
	}
	// the smallest unit like lamports
	fmt.Println("balance", balance)
	// the decimals of mint which token account holds
	fmt.Println("decimals", decimals)

	// if you want use a normal unit, you can do
	// balance / 10^decimals
}
