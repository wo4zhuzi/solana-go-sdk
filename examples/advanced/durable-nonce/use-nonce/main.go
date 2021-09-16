package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wo4zhuzi/solana-go-sdk/client"
	"github.com/wo4zhuzi/solana-go-sdk/common"
	"github.com/wo4zhuzi/solana-go-sdk/program/sysprog"
	"github.com/wo4zhuzi/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBytes([]byte{178, 244, 76, 4, 247, 41, 113, 40, 111, 103, 12, 76, 195, 4, 100, 123, 88, 226, 37, 56, 209, 180, 92, 77, 39, 85, 78, 202, 121, 162, 88, 29, 125, 155, 223, 107, 139, 223, 229, 82, 89, 209, 27, 43, 108, 205, 144, 2, 74, 159, 215, 57, 198, 4, 193, 36, 161, 50, 160, 119, 89, 240, 102, 184})

func main() {
	// you created before
	nonceAccountPubkey := common.PublicKeyFromString("CjJWxNi3j8PyxSuTwSiJnLSbKuzV5JgRi8WpdPz1LzPX")

	c := client.NewClient("http://localhost:8899")

	// fetch nonce
	accountInfo, err := c.GetAccountInfo(
		context.Background(),
		"CjJWxNi3j8PyxSuTwSiJnLSbKuzV5JgRi8WpdPz1LzPX",
	)
	if err != nil {
		log.Fatalf("failed to get account info, err: %v", err)
	}
	nonceAccount, err := sysprog.NonceAccountDeserialize(accountInfo.Data)
	if err != nil {
		log.Fatalf("failed to deserialize nonce account, err: %v", err)
	}

	// create a random account
	to := types.NewAccount()

	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			// you need to put this instruction in the first
			sysprog.AdvanceNonceAccount(
				nonceAccountPubkey,
				feePayer.PublicKey, // nonce account's owner
			),
			// now put the instrucitons you really want to do
			// here I use transfer as a example
			sysprog.Transfer(
				feePayer.PublicKey,
				to.PublicKey,
				1e9,
			),
		},
		Signers:  []types.Account{feePayer},
		FeePayer: feePayer.PublicKey,
		// here must use the nonce account's nonce as a recent blockhash
		RecentBlockHash: nonceAccount.Nonce.ToBase58(),
	})
	if err != nil {
		log.Fatalf("failed to create raw transaction, err: %v", err)
	}

	sig, err := c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("failed to send transaction, err: %v", err)
	}

	fmt.Println(sig)
}
