package rpc

import (
	"context"
	"fmt"
)

type GetBlockConfig struct {
	// TODO custom
	Encoding string `json:"encoding"` // default: "json", either "json", "jsonParsed", "base58" (slow), "base64"
	//TransactionDetails string     `json:"transactionDetails"`   // default: "full", either "full", "signatures", "none"
	MaxSupportedTransactionVersion int        `json:"maxSupportedTransactionVersion"` //  set the max transaction version to return in responses. If the requested block contains a transaction with a higher version, an error will be returned. If this parameter is omitted, only legacy transactions will be returned, and a block containing any versioned transaction will prompt the error.
	Commitment                     Commitment `json:"commitment,omitempty"`           // "processed" is not supported. If parameter not provided, the default is "finalized".
}

type GetBlockResponse struct {
	BlockHeight       int64  `json:"blockHeight"`
	Blockhash         string `json:"blockhash"`
	PreviousBlockhash string `json:"previousBlockhash"`
	ParentSLot        uint64 `json:"parentSlot"`
	BlockTime         int64  `json:"blockTime"`
	Transactions      []struct {
		Meta        TransactionMetaJsonParsed `json:"meta"`
		Transaction Transaction               `json:"transaction"`
	} `json:"transactions"`
	Rewards []struct {
		Pubkey      string `json:"pubkey"`
		Lamports    int64  `json:"lamports"`
		PostBalance uint64 `json:"postBalance"`
		RewardType  string `json:"rewardType"` // type of reward: "fee", "rent", "voting", "staking"
	} `json:"rewards"`
}

// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlock for solana-core v1.6
// GetBlock returns identity and transaction information about a confirmed block in the ledger
func (s *RpcClient) GetBlock(ctx context.Context, slot uint64, cfg GetBlockConfig) (GetBlockResponse, error) {
	res := struct {
		GeneralResponse
		Result GetBlockResponse `json:"result"`
	}{}
	err := s.request(ctx, "getBlock", []interface{}{slot, cfg}, &res)
	if err != nil {
		return GetBlockResponse{}, err
	}

	if res.Error != nil && res.Error.Code != -32007 && res.Error.Code != -32009 && res.Error.Code != -32015 {
		return GetBlockResponse{}, fmt.Errorf("%v", res.Error)
	}

	return res.Result, nil
}
