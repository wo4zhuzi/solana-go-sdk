package rpc

import (
	"context"
)

type GetTokenAccountsByOwnerResponse struct {
	GeneralResponse
	Result struct {
		Context struct {
			Slot int `json:"slot"`
		} `json:"context"`
		Value GetTokenAccountsByOwnerResponseValue `json:"value"`
	} `json:"result"`
}

type GetTokenAccountsByOwnerResponseValue []struct {
	Account struct {
		Data struct {
			Parsed struct {
				Info struct {
					IsNative bool `json:"isNative"`
					Mint string `json:"mint"`
					Owner string `json:"owner"`
					State string `json:"state"`
					TokenAmount struct {
						Amount string `json:"amount"`
						Decimals int `json:"decimals"`
						UIAmount float64 `json:"uiAmount"`
						UIAmountString string `json:"uiAmountString"`
					} `json:"tokenAmount"`
				} `json:"info"`
				Type string `json:"type"`
			} `json:"parsed"`
			Program string `json:"program"`
			Space int `json:"space"`
		} `json:"data"`
		Executable bool `json:"executable"`
		Lamports int `json:"lamports"`
		Owner string `json:"owner"`
		RentEpoch int `json:"rentEpoch"`
	} `json:"account"`
	Pubkey string `json:"pubkey"`
}

func (s *RpcClient) GetTokenAccountsByOwner(ctx context.Context, address string, mint string) (GetTokenAccountsByOwnerResponse, error) {
	res := GetTokenAccountsByOwnerResponse{}
	err := s.request(ctx, "getTokenAccountsByOwner", []interface{}{address, map[string]string{"mint" : mint}, map[string]string{"encoding" : "jsonParsed"}}, &res)
	if err != nil {
		return GetTokenAccountsByOwnerResponse{}, err
	}

	return res, nil
}
