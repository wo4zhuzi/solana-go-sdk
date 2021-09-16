package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wo4zhuzi/solana-go-sdk/client"
	"github.com/wo4zhuzi/solana-go-sdk/client/rpc"
	"github.com/wo4zhuzi/solana-go-sdk/types"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	newAccount := types.NewAccount()
	fmt.Println(newAccount.PublicKey.ToBase58())

	txhash, err := c.RequestAirdrop(
		context.Background(),
		newAccount.PublicKey.ToBase58(),
		1e9, // 1 SOL = 10^9 lamports
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}

	fmt.Println("txhash:", txhash)
}
