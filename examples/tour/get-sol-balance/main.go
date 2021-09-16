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
	balance, err := c.GetBalance(
		context.Background(),
		"RNfp4xTbBb4C3kcv2KqtAj8mu4YhMHxqm1Skg9uchZ7",
	)
	if err != nil {
		log.Fatalln("get balance error", err)
	}
	fmt.Println(balance)
}
